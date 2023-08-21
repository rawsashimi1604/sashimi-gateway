package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/route"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rs/zerolog/log"
)

func NewPostgresServiceGateway(conn *pgxpool.Pool) *PostgresServiceGateway {
	return &PostgresServiceGateway{Conn: conn}
}

func (s *PostgresServiceGateway) GetServiceByPath(path string) (models.Service, error) {
	query := `
		SELECT s.id, s.name, s.target_url, s.path, s.description, s.created_at, s.updated_at, r.id, r.path, r.description, r.created_at, r.updated_at
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
	var routes = make([]route.Route_DB, 0)
	serviceExists := false
	for rows.Next() {
		serviceExists = true
		var route route.Route_DB
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
		); err != nil {
			log.Info().Msg(err.Error())
			return models.Service{}, err
		}
		routes = append(routes, route)
	}

	if !serviceExists {
		return models.Service{}, ErrServiceNotFound
	}

	return MapServiceDbToDomain(service, routes), nil
}

func (s *PostgresServiceGateway) GetServiceByTargetUrl(targetUrl string) (models.Service, error) {
	query := `
		SELECT s.id, s.name, s.target_url, s.path, s.description, s.created_at, s.updated_at, r.id, r.path, r.description, r.created_at, r.updated_at
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
	var routes = make([]route.Route_DB, 0)
	serviceExists := false

	for rows.Next() {
		serviceExists = true
		var route route.Route_DB
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
		); err != nil {
			log.Info().Msg(err.Error())
			return models.Service{}, err
		}
		routes = append(routes, route)
	}

	if !serviceExists {
		return models.Service{}, ErrServiceNotFound
	}

	return MapServiceDbToDomain(service, routes), nil
}

func (s *PostgresServiceGateway) GetAllServices() ([]models.Service, error) {
	query := `
		SELECT s.id, s.name, s.target_url, s.path, s.description, s.created_at, s.updated_at, r.id, r.path, r.description, r.created_at, r.updated_at
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
	serviceRoutesMap := make(map[int][]route.Route_DB)
	serviceMap := make(map[int]Service_DB)

	for rows.Next() {
		var service Service_DB
		var r route.Route_DB

		if err := rows.Scan(
			&service.Id,
			&service.Name,
			&service.TargetUrl,
			&service.Path,
			&service.Description,
			&service.CreatedAt,
			&service.UpdatedAt,
			&r.Id,
			&r.Path,
			&r.Description,
			&r.CreatedAt,
			&r.UpdatedAt,
		); err != nil {
			log.Info().Msg("error retrieving service and route.")
			continue
		}

		// If this service is not in the serviceMap, add it
		if _, exists := serviceMap[service.Id]; !exists {
			serviceMap[service.Id] = service
		}

		// Append the route to the service's routes in the map
		serviceRoutesMap[service.Id] = append(serviceRoutesMap[service.Id], r)
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
