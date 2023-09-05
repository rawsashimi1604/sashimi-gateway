package request

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
)

type RequestGateway interface {
	AddBulkRequests(requests []models.ApiRequest) ([]models.ApiRequest, error)
	GetAllRequests() ([]models.ApiRequest, error)
}

type PostgresRequestGateway struct {
	Conn *pgxpool.Pool
}

type Request_DB struct {
	Id        string    `json:"id"`
	ServiceId int       `json:"serviceId"`
	RouteId   int       `json:"routeId"`
	Path      string    `json:"path"`
	Method    string    `json:"method"`
	Time      time.Time `json:"time"`
	Code      int       `json:"code"`
}

func MapRequestDbToDomain(rdb Request_DB) models.ApiRequest {
	id, _ := uuid.Parse(rdb.Id)

	return models.ApiRequest{
		Id:        id,
		ServiceId: rdb.ServiceId,
		RouteId:   rdb.RouteId,
		Path:      rdb.Path,
		Method:    rdb.Method,
		Time:      rdb.Time,
		Code:      rdb.Code,
	}
}
