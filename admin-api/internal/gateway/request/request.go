package request

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
)

type RequestGateway interface {
	AddBulkRequests([]models.ApiRequest) ([]models.ApiRequest, error)
}

type PostgresRequestGateway struct {
	Conn *pgxpool.Pool
}

type Request_DB struct {
	Id        int    `json:"id"`
	ServiceId int    `json:"serviceId"`
	RouteId   int    `json:"routeId"`
	Path      string `json:"path"`
	MethodId  string `json:"methodId"`
	Time      time.Time
}
