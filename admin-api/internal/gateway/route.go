package gateway

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rs/zerolog/log"
)

type RouteGateway struct {
	Conn *pgxpool.Pool
}

type Route_DB struct {
	Id          int       `json:"id"`
	Path        string    `json:"path"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

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

func NewRouteGateway(conn *pgxpool.Pool) *RouteGateway {
	return &RouteGateway{Conn: conn}
}

func (s *RouteGateway) GetAllRoutes() ([]models.Route, error) {

	query := `
		SELECT id, path, description, created_at, updated_at
		FROM route
		ORDER BY id ASC
	`

	rows, err := s.Conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var routes []Route_DB
	for rows.Next() {
		var route Route_DB
		if err := rows.Scan(&route.Id, &route.Path, &route.Description, &route.CreatedAt, &route.UpdatedAt); err != nil {
			log.Info().Msg("error retrieving route.")
			continue
		}
		routes = append(routes, route)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	var mappedDbs []models.Route
	for _, routeDb := range routes {
		mapped := mapRouteDbToDomain(routeDb)
		mappedDbs = append(mappedDbs, mapped)
	}
	return mappedDbs, nil
}
