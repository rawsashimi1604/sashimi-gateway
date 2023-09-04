package request

import (
	"bytes"
	"context"
	"fmt"

	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/utils"
	"github.com/rs/zerolog/log"
)

func NewPostgresRequestGatweay(conn *pgxpool.Pool) *PostgresRequestGateway {
	return &PostgresRequestGateway{Conn: conn}
}

func (rg *PostgresRequestGateway) AddBulkRequests(requests []models.ApiRequest) ([]models.ApiRequest, error) {
	// Bulk insert using COPY command from Buffer
	var byteBuffer bytes.Buffer
	for _, request := range requests {
		fmt.Fprintf(&byteBuffer, "%d\t%d\t%d\t%s\t%s\t%s\n", request.Id, request.ServiceId, request.RouteId, request.Path, request.Method, request.Time.Format(time.RFC3339))
	}

	copyBuffer := utils.NewCopyBuffer(&byteBuffer)
	_, err := rg.Conn.CopyFrom(context.Background(),
		pgx.Identifier{"api_request"},
		[]string{"id", "service_id", "route_id", "path", "method", "time"},
		copyBuffer,
	)

	if err != nil {
		log.Info().Msg("something went wrong when adding bulk requests")
	}

}
