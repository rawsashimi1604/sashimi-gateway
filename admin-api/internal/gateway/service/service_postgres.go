package service

import (
	"context"

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
		SELECT s.id, s.name, s.target_url, s.path, s.description, s.created_at, s.updated_at, r.id, r.path, r.description, r.created_at, r.updated_at, am.id, am.method
		FROM service s
		LEFT JOIN route r
		ON s.id=r.service_id
		INNER JOIN api_method am
		ON r.method_id=am.id
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
		var apiMethod r.ApiMethod_DB

		if err := rows.Scan(
			&service.Id,
			&service.Name,
			&service.TargetUrl,
			&service.Path,
			&service.Description,
			&service.CreatedAt,
			&service.UpdatedAt,
			&route.Id,
			&route.Path,
			&route.Description,
			&route.CreatedAt,
			&route.UpdatedAt,
			&apiMethod.Id,
			&apiMethod.Method,
		); err != nil {
			log.Info().Msg(err.Error())
			return models.Service{}, err
		}

		routes = append(routes, r.MapRouteDbToDomain(route, apiMethod))
	}

	if !serviceExists {
		return models.Service{}, ErrServiceNotFound
	}

	return MapServiceDbToDomain(service, routes), nil
}

func (s *PostgresServiceGateway) GetServiceByTargetUrl(targetUrl string) (models.Service, error) {
	query := `
		SELECT s.id, s.name, s.target_url, s.path, s.description, s.created_at, s.updated_at, r.id, r.path, r.description, r.created_at, r.updated_at, am.id, am.method
		FROM service s
		LEFT JOIN route r
		ON s.id=r.service_id
		INNER JOIN api_method am
		ON r.method_id=am.id
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
		var apiMethod r.ApiMethod_DB

		if err := rows.Scan(
			&service.Id,
			&service.Name,
			&service.TargetUrl,
			&service.Path,
			&service.Description,
			&service.CreatedAt,
			&service.UpdatedAt,
			&route.Id,
			&route.Path,
			&route.Description,
			&route.CreatedAt,
			&route.UpdatedAt,
			&apiMethod.Id,
			&apiMethod.Method,
		); err != nil {
			log.Info().Msg(err.Error())
			return models.Service{}, err
		}

		routes = append(routes, r.MapRouteDbToDomain(route, apiMethod))
	}

	if !serviceExists {
		return models.Service{}, ErrServiceNotFound
	}

	return MapServiceDbToDomain(service, routes), nil
}

func (s *PostgresServiceGateway) GetAllServices() ([]models.Service, error) {
	query := `
		SELECT s.id, s.name, s.target_url, s.path, s.description, s.created_at, s.updated_at, r.id, r.path, r.description, r.created_at, r.updated_at, am.id, am.method
		FROM service s
		LEFT JOIN route r
		ON s.id=r.service_id
		INNER JOIN api_method am
		ON r.method_id=am.id
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
		var apiMethod r.ApiMethod_DB

		if err := rows.Scan(
			&service.Id,
			&service.Name,
			&service.TargetUrl,
			&service.Path,
			&service.Description,
			&service.CreatedAt,
			&service.UpdatedAt,
			&route.Id,
			&route.Path,
			&route.Description,
			&route.CreatedAt,
			&route.UpdatedAt,
			&apiMethod.Id,
			&apiMethod.Method,
		); err != nil {
			log.Info().Msg("error retrieving service and route.")
			continue
		}

		// If this service is not in the serviceMap, add it
		if _, exists := serviceMap[service.Id]; !exists {
			serviceMap[service.Id] = service
		}

		// Append the route to the service's routes in the map
		serviceRoutesMap[service.Id] = append(serviceRoutesMap[service.Id], r.MapRouteDbToDomain(route, apiMethod))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	var mappedDbs []models.Service
	for _, serviceDb := range serviceMap {
		routes := serviceRoutesMap[serviceDb.Id]
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

	createdService := models.Service{
		Routes: make([]models.Route, 0),
	}

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

	return createdService, nil
}
