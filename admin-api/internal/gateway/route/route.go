package route

import (
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
)

type RouteGateway interface {
	GetAllRoutes() ([]models.Route, error)
	RegisterRoute(route models.Route) (models.Route, error)
}

type PostgresRouteGateway struct {
	Conn *pgxpool.Pool
}

type Route_DB struct {
	Id          int       `json:"id"`
	ServiceId   int       `json:"serviceId"`
	Path        string    `json:"path"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Method      string    `json:"method"`
}

var (
	ErrRouteNotFound = errors.New("route not found in the database")
)

func MapRouteDbToDomain(rdb Route_DB) models.Route {
	return models.Route{
		Id:          rdb.Id,
		ServiceId:   rdb.ServiceId,
		Path:        rdb.Path,
		Description: rdb.Description,
		CreatedAt:   rdb.CreatedAt,
		UpdatedAt:   rdb.UpdatedAt,
		Method:      rdb.Method,
	}
}

func NewRouteGateway(conn *pgxpool.Pool) *PostgresRouteGateway {
	return &PostgresRouteGateway{Conn: conn}
}
