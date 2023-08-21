package service

import (
	"context"

	"github.com/jackc/pgx/v5"
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
		INNER JOIN route r
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

	if err := rows.Err(); err != nil {
		return models.Service{}, err
	}

	return MapServiceDbToDomain(service, routes), nil
}

func (s *PostgresServiceGateway) GetServiceByTargetUrl(targetUrl string) (models.Service, error) {
	query := `
		SELECT id, name, target_url, path, description, created_at, updated_at
		FROM service
		WHERE target_url=$1
	`

	var service Service_DB
	if err := s.Conn.QueryRow(context.Background(), query, targetUrl).Scan(
		&service.Id,
		&service.Name,
		&service.TargetUrl,
		&service.Path,
		&service.Description,
		&service.CreatedAt,
		&service.UpdatedAt,
	); err != nil {
		if err == pgx.ErrNoRows {
			return models.Service{}, ErrServiceNotFound
		}
		return models.Service{}, err
	}

	return MapServiceDbToDomain(service, make([]route.Route_DB, 0)), nil
}

func (s *PostgresServiceGateway) GetAllServices() ([]models.Service, error) {

	query := `
		SELECT id, name, target_url, path, description, created_at, updated_at
		FROM service
		ORDER BY id ASC
	`

	rows, err := s.Conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []Service_DB
	for rows.Next() {
		var service Service_DB
		if err := rows.Scan(&service.Id, &service.Name, &service.TargetUrl, &service.Path, &service.Description, &service.CreatedAt, &service.UpdatedAt); err != nil {
			log.Info().Msg("error retrieving service.")
			continue
		}
		services = append(services, service)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	var mappedDbs []models.Service
	for _, serviceDb := range services {
		mapped := MapServiceDbToDomain(serviceDb, make([]route.Route_DB, 0))
		mappedDbs = append(mappedDbs, mapped)
	}
	return mappedDbs, nil
}
