package route

import (
	"context"

	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rs/zerolog/log"
)

func (s *PostgresRouteGateway) GetAllRoutes() ([]models.Route, error) {

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
