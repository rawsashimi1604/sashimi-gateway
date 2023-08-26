package route

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
)

func NewPostgresRouteGateway(conn *pgxpool.Pool) *PostgresRouteGateway {
	return &PostgresRouteGateway{Conn: conn}
}

func (s *PostgresRouteGateway) GetAllRoutes() ([]models.Route, error) {

	query := `
		SELECT r.id, r.path, r.description, r.created_at, r.updated_at, m.id, m.method
		FROM route r
		LEFT JOIN api_method m
		ON r.method_id=m.id
		ORDER BY r.id ASC
	`

	rows, err := s.Conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var routes []models.Route
	for rows.Next() {
		var route Route_DB
		var method ApiMethod_DB

		if err := rows.Scan(&route.Id, &route.Path, &route.Description, &route.CreatedAt, &route.UpdatedAt, &method.Id, &method.Method); err != nil {
			return nil, errors.New("error retrieving route")
		}

		routes = append(routes, MapRouteDbToDomain(route, method))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return routes, nil
}
