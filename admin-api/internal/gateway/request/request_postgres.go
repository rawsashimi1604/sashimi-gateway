package request

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/utils"
	"github.com/rs/zerolog/log"
)

func NewPostgresRequestGateway(conn *pgxpool.Pool) *PostgresRequestGateway {
	return &PostgresRequestGateway{Conn: conn}
}

func (rg *PostgresRequestGateway) AddBulkRequests(requests []models.ApiRequest) ([]models.ApiRequest, error) {
	// Bulk insert using COPY command from Buffer
	var byteBuffer bytes.Buffer
	for _, request := range requests {
		fmt.Fprintf(
			&byteBuffer,
			"%s\t%d\t%d\t%s\t%s\t%s\t%d\t%d\n",
			request.Id,
			request.ServiceId,
			request.RouteId,
			request.Path,
			request.Method,
			request.Time.UTC().Format("2006-01-02 15:04:05"),
			request.Code,
			request.Duration,
		)
	}

	copyBuffer := utils.NewCopyBuffer(&byteBuffer)
	_, err := rg.Conn.CopyFrom(context.Background(),
		pgx.Identifier{"api_request"},
		[]string{"id", "service_id", "route_id", "path", "method", "time", "code", "duration"},
		copyBuffer,
	)

	if err != nil {
		log.Info().Msg("something went wrong when adding bulk requests")
		return nil, err
	}

	return requests, nil
}

func (rg *PostgresRequestGateway) GetAllRequests() ([]models.ApiRequest, error) {
	query := `
	SELECT id, service_id, route_id, path, method, time, code, duration
	FROM api_request
`

	rows, err := rg.Conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apiRequests []models.ApiRequest
	for rows.Next() {
		var apiRequest Request_DB

		if err := rows.Scan(
			&apiRequest.Id,
			&apiRequest.ServiceId,
			&apiRequest.RouteId,
			&apiRequest.Path,
			&apiRequest.Method,
			&apiRequest.Time,
			&apiRequest.Code,
			&apiRequest.Duration,
		); err != nil {
			return nil, errors.New("error retrieving api requests")
		}

		apiRequests = append(apiRequests, MapRequestDbToDomain(apiRequest))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return apiRequests, nil
}

func (rg *PostgresRequestGateway) GetAggregatedRequests(timespan int, dataPoints int) ([]models.AggregatedApiRequest, error) {
	query := `
	WITH time_series AS (
		SELECT 
			start_bin, 
			start_bin + $1::integer * INTERVAL '1 MINUTE' AS end_bin
		FROM generate_series(
			(NOW() AT TIME ZONE 'UTC' - ($2::integer * $1::integer * INTERVAL '1 MINUTE')),
			(NOW() AT TIME ZONE 'UTC'),
			$1::integer * INTERVAL '1 MINUTE'
		) AS t(start_bin)
	)
		
	SELECT
		ts.start_bin,
		COALESCE(COUNT(ar.id), 0) AS count,
		COALESCE(COUNT(CASE WHEN ar.code >= 200 AND ar.code < 300 THEN ar.id END), 0) AS count_2xx,
		COALESCE(COUNT(CASE WHEN ar.code >= 400 AND ar.code < 500 THEN ar.id END), 0) AS count_4xx,
		COALESCE(COUNT(CASE WHEN ar.code >= 500 AND ar.code < 600 THEN ar.id END), 0) AS count_5xx
	FROM 
		time_series ts 
	LEFT JOIN 
		api_request ar ON ar.time >= ts.start_bin AND ar.time < ts.end_bin
	GROUP BY 
		ts.start_bin
	ORDER BY 
		ts.start_bin;
	
	`

	rows, err := rg.Conn.Query(context.Background(), query, timespan, dataPoints)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	allAggregatedBuckets := make([]models.AggregatedApiRequest, 0)

	for rows.Next() {
		var aggregated models.AggregatedApiRequest

		if err := rows.Scan(&aggregated.TimeBucket, &aggregated.Count, &aggregated.Count2xx, &aggregated.Count4xx, &aggregated.Count5xx); err != nil {
			log.Info().Msg(err.Error())
			return nil, errors.New("error retrieving aggregated api requests")
		}
		allAggregatedBuckets = append(allAggregatedBuckets, aggregated)
	}

	return allAggregatedBuckets, nil
}
