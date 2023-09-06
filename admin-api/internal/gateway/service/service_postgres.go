package service

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
	r "github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/route"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rs/zerolog/log"
)

func NewPostgresServiceGateway(conn *pgxpool.Pool) *PostgresServiceGateway {
	return &PostgresServiceGateway{Conn: conn}
}

func (s *PostgresServiceGateway) GetServiceByPath(path string) (models.Service, error) {
	query := `
		SELECT s.id, s.name, s.target_url, s.path, s.description, s.created_at, s.updated_at, r.id, r.path, r.description, r.created_at, r.updated_at, r.method
		FROM service s
		LEFT JOIN route r
		ON s.id=r.service_id
		WHERE s.path=$1
	`

	rows, err := s.Conn.Query(context.Background(), query, path)
	if err != nil {
		return models.Service{}, err
	}
	defer rows.Close()

	var service Service_DB
	var routes = make([]models.Route, 0)

	serviceExists := false
	for rows.Next() {
		serviceExists = true
		var route r.Route_DB

		var routeID sql.NullInt64
		var routePath sql.NullString
		var routeDescription sql.NullString
		var routeCreatedAt, routeUpdatedAt sql.NullTime
		var routeMethod sql.NullString

		if err := rows.Scan(
			&service.Id,
			&service.Name,
			&service.TargetUrl,
			&service.Path,
			&service.Description,
			&service.CreatedAt,
			&service.UpdatedAt,
			&routeID,
			&routePath,
			&routeDescription,
			&routeCreatedAt,
			&routeUpdatedAt,
			&routeMethod,
		); err != nil {
			log.Info().Msg(err.Error())
			return models.Service{}, err
		}

		if routeID.Valid {
			route.Id = int(routeID.Int64)
			route.Path = routePath.String
			route.Description = routeDescription.String
			route.CreatedAt = routeCreatedAt.Time
			route.UpdatedAt = routeUpdatedAt.Time
			route.Method = routeMethod.String
			routes = append(routes, r.MapRouteDbToDomain(route))
		}
	}

	if !serviceExists {
		return models.Service{}, ErrServiceNotFound
	}

	return MapServiceDbToDomain(service, routes), nil
}

func (s *PostgresServiceGateway) GetServiceByTargetUrl(targetUrl string) (models.Service, error) {
	query := `
		SELECT s.id, s.name, s.target_url, s.path, s.description, s.created_at, s.updated_at, r.id, r.path, r.description, r.created_at, r.updated_at, r.method
		FROM service s
		LEFT JOIN route r
		ON s.id=r.service_id
		WHERE s.target_url=$1
	`

	rows, err := s.Conn.Query(context.Background(), query, targetUrl)
	if err != nil {
		return models.Service{}, err
	}
	defer rows.Close()

	var service Service_DB
	var routes = make([]models.Route, 0)
	serviceExists := false

	for rows.Next() {
		serviceExists = true

		var route r.Route_DB

		var routeID sql.NullInt64
		var routePath sql.NullString
		var routeDescription sql.NullString
		var routeCreatedAt, routeUpdatedAt sql.NullTime
		var routeMethod sql.NullString

		if err := rows.Scan(
			&service.Id,
			&service.Name,
			&service.TargetUrl,
			&service.Path,
			&service.Description,
			&service.CreatedAt,
			&service.UpdatedAt,
			&routeID,
			&routePath,
			&routeDescription,
			&routeCreatedAt,
			&routeUpdatedAt,
			&routeMethod,
		); err != nil {
			log.Info().Msg(err.Error())
			return models.Service{}, err
		}

		if routeID.Valid {
			route.Id = int(routeID.Int64)
			route.ServiceId = service.Id
			route.Path = routePath.String
			route.Description = routeDescription.String
			route.CreatedAt = routeCreatedAt.Time
			route.UpdatedAt = routeUpdatedAt.Time
			route.Method = routeMethod.String
			routes = append(routes, r.MapRouteDbToDomain(route))
		}
	}

	if !serviceExists {
		return models.Service{}, ErrServiceNotFound
	}

	return MapServiceDbToDomain(service, routes), nil
}

func (s *PostgresServiceGateway) GetServiceById(id int) (models.Service, error) {
	query := `
		SELECT s.id, s.name, s.target_url, s.path, s.description, s.created_at, s.updated_at, r.id, r.path, r.description, r.created_at, r.updated_at, r.method
		FROM service s
		LEFT JOIN route r
		ON s.id=r.service_id
		WHERE s.id=$1
	`

	rows, err := s.Conn.Query(context.Background(), query, id)
	if err != nil {
		return models.Service{}, err
	}
	defer rows.Close()

	var service Service_DB
	var routes = make([]models.Route, 0)

	serviceExists := false
	for rows.Next() {
		serviceExists = true
		var route r.Route_DB

		var routeID sql.NullInt64
		var routePath sql.NullString
		var routeDescription sql.NullString
		var routeCreatedAt, routeUpdatedAt sql.NullTime
		var routeMethod sql.NullString

		if err := rows.Scan(
			&service.Id,
			&service.Name,
			&service.TargetUrl,
			&service.Path,
			&service.Description,
			&service.CreatedAt,
			&service.UpdatedAt,
			&routeID,
			&routePath,
			&routeDescription,
			&routeCreatedAt,
			&routeUpdatedAt,
			&routeMethod,
		); err != nil {
			log.Info().Msg(err.Error())
			return models.Service{}, err
		}

		if routeID.Valid {
			route.Id = int(routeID.Int64)
			route.Path = routePath.String
			route.Description = routeDescription.String
			route.CreatedAt = routeCreatedAt.Time
			route.UpdatedAt = routeUpdatedAt.Time
			route.Method = routeMethod.String
			routes = append(routes, r.MapRouteDbToDomain(route))
		}
	}

	if !serviceExists {
		return models.Service{}, ErrServiceNotFound
	}

	return MapServiceDbToDomain(service, routes), nil
}

func (s *PostgresServiceGateway) GetAllServices() ([]models.Service, error) {
	query := `
		SELECT s.id, s.name, s.target_url, s.path, s.description, s.created_at, s.updated_at, r.id, r.path, r.description, r.created_at, r.updated_at, r.method
		FROM service s
		LEFT JOIN route r
		ON s.id=r.service_id
		ORDER BY s.id
	`

	rows, err := s.Conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Service id -> slice of Route_DB
	serviceRoutesMap := make(map[int][]models.Route)
	serviceMap := make(map[int]Service_DB)

	for rows.Next() {
		var service Service_DB
		var route r.Route_DB

		var routeID sql.NullInt64
		var routePath sql.NullString
		var routeDescription sql.NullString
		var routeCreatedAt, routeUpdatedAt sql.NullTime
		var routeMethod sql.NullString

		if err := rows.Scan(
			&service.Id,
			&service.Name,
			&service.TargetUrl,
			&service.Path,
			&service.Description,
			&service.CreatedAt,
			&service.UpdatedAt,
			&routeID,
			&routePath,
			&routeDescription,
			&routeCreatedAt,
			&routeUpdatedAt,
			&routeMethod,
		); err != nil {
			log.Info().Msg("error retrieving service and route.")
			continue
		}

		// If this service is not in the serviceMap, add it
		if _, exists := serviceMap[service.Id]; !exists {
			serviceMap[service.Id] = service
		}

		if routeID.Valid {
			route.Id = int(routeID.Int64)
			route.ServiceId = service.Id
			route.Path = routePath.String
			route.Description = routeDescription.String
			route.CreatedAt = routeCreatedAt.Time
			route.UpdatedAt = routeUpdatedAt.Time
			route.Method = routeMethod.String

			// Append the route to the service's routes in the map
			serviceRoutesMap[service.Id] = append(serviceRoutesMap[service.Id], r.MapRouteDbToDomain(route))
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	var mappedDbs []models.Service
	for _, serviceDb := range serviceMap {
		routes := serviceRoutesMap[serviceDb.Id]
		if routes == nil {
			routes = make([]models.Route, 0)
		}
		mapped := MapServiceDbToDomain(serviceDb, routes)
		mappedDbs = append(mappedDbs, mapped)
	}

	return mappedDbs, nil
}

func (s *PostgresServiceGateway) RegisterService(service models.Service) (models.Service, error) {
	// To be completed.
	query := `
		INSERT INTO service 
			(name, target_url, path, description, created_at, updated_at) 
		VALUES
			($1, $2, $3, $4, $5, $6)
		RETURNING id, name, target_url, path, description, created_at, updated_at;
	`

	row := s.Conn.QueryRow(
		context.Background(),
		query,
		service.Name,
		service.TargetUrl,
		service.Path,
		service.Description,
		service.CreatedAt,
		service.UpdatedAt,
	)

	createdService := Service_DB{}

	if err := row.Scan(
		&createdService.Id,
		&createdService.Name,
		&createdService.TargetUrl,
		&createdService.Path,
		&createdService.Description,
		&createdService.CreatedAt,
		&createdService.UpdatedAt,
	); err != nil {
		log.Info().Msg(err.Error())
		return models.Service{}, err
	}

	return MapServiceDbToDomain(createdService, make([]models.Route, 0)), nil
}
