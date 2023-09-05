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
			"%s\t%d\t%d\t%s\t%s\t%s\t%d\n",
			request.Id,
			request.ServiceId,
			request.RouteId,
			request.Path,
			request.Method,
			request.Time.UTC().Format("2006-01-02 15:04:05"),
			request.Code,
		)
	}

	copyBuffer := utils.NewCopyBuffer(&byteBuffer)
	_, err := rg.Conn.CopyFrom(context.Background(),
		pgx.Identifier{"api_request"},
		[]string{"id", "service_id", "route_id", "path", "method", "time", "code"},
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
	SELECT id, service_id, route_id, path, method, time, code
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
