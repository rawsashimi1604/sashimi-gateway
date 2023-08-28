package route

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rs/zerolog/log"
)

func NewPostgresRouteGateway(conn *pgxpool.Pool) *PostgresRouteGateway {
	return &PostgresRouteGateway{Conn: conn}
}

func (rg *PostgresRouteGateway) GetAllRoutes() ([]models.Route, error) {

	query := `
		SELECT r.id, r.service_id, r.path, r.description, r.created_at, r.updated_at, m.id, m.method
		FROM route r
		LEFT JOIN api_method m
		ON r.method_id=m.id
		ORDER BY r.id ASC
	`

	rows, err := rg.Conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var routes []models.Route
	for rows.Next() {
		var route Route_DB
		var method ApiMethod_DB

		if err := rows.Scan(&route.Id, &route.ServiceId, &route.Path, &route.Description, &route.CreatedAt, &route.UpdatedAt, &method.Id, &method.Method); err != nil {
			return nil, errors.New("error retrieving route")
		}

		routes = append(routes, MapRouteDbToDomain(route, method))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return routes, nil
}

func (rg *PostgresRouteGateway) RegisterRoute(route models.Route) (models.Route, error) {
	query := `
	WITH inserted AS (
		INSERT INTO route
			(service_id, method_id, path, description, created_at, updated_at)
		VALUES 
			($1, $2, $3, $4, $5, $6)
		RETURNING *
	)
	SELECT inserted.id, service_id, method_id, method, path, description, created_at, updated_at
	FROM inserted
	INNER JOIN api_method ON inserted.method_id = api_method.id
	`

	row := rg.Conn.QueryRow(
		context.Background(),
		query,
		route.ServiceId,
		route.Method.Id,
		route.Path,
		route.Description,
		route.CreatedAt,
		route.UpdatedAt,
	)

	createdRoute := Route_DB{}
	apiMethodDb := ApiMethod_DB{}

	if err := row.Scan(
		&createdRoute.Id,
		&createdRoute.ServiceId,
		&apiMethodDb.Id,
		&apiMethodDb.Method,
		&createdRoute.Path,
		&createdRoute.Description,
		&createdRoute.CreatedAt,
		&createdRoute.UpdatedAt,
	); err != nil {
		log.Info().Msg(err.Error())
		return models.Route{}, err
	}

	return MapRouteDbToDomain(createdRoute, apiMethodDb), nil
}
