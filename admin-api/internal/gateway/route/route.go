package route

import (
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
)

type RouteGateway interface {
	GetAllRoutes() ([]models.Route, error)
}

type PostgresRouteGateway struct {
	Conn *pgxpool.Pool
}

type Route_DB struct {
	Id          int       `json:"id"`
	Path        string    `json:"path"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var (
	ErrRouteNotFound = errors.New("route not found in the database")
)

func mapRouteDbToDomain(rdb Route_DB) models.Route {
	return models.Route{
		Id:          rdb.Id,
		Path:        rdb.Path,
		Description: rdb.Description,
		CreatedAt:   rdb.CreatedAt,
		UpdatedAt:   rdb.UpdatedAt,
		Methods:     make([]models.ApiMethod, 0),
	}
}

func NewRouteGateway(conn *pgxpool.Pool) *PostgresRouteGateway {
	return &PostgresRouteGateway{Conn: conn}
}
