package gateway

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rs/zerolog/log"
)

type ServiceGateway struct {
	Conn *pgxpool.Pool
}

type Service_DB struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	TargetUrl   string    `json:"targetUrl"`
	Path        string    `json:"path"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var (
	ErrServiceNotFound = errors.New("service not found in the database")
)

func mapServiceDbToDomain(sdb Service_DB) models.Service {
	return models.Service{
		Id:          sdb.Id,
		Name:        sdb.Name,
		TargetUrl:   sdb.TargetUrl,
		Path:        sdb.Path,
		Description: sdb.Description,
		CreatedAt:   sdb.CreatedAt,
		UpdatedAt:   sdb.UpdatedAt,
		Routes:      make([]models.Route, 0),
	}
}

func NewServiceGateway(conn *pgxpool.Pool) *ServiceGateway {
	return &ServiceGateway{Conn: conn}
}

func (s *ServiceGateway) GetServiceByPath(path string) (models.Service, error) {
	query := `
		SELECT id, name, target_url, path, description, created_at, updated_at
		FROM service
		WHERE path=$1
	`

	var service Service_DB
	if err := s.Conn.QueryRow(context.Background(), query, path).Scan(
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

	return mapServiceDbToDomain(service), nil
}

func (s *ServiceGateway) GetServiceByTargetUrl(targetUrl string) (models.Service, error) {
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

	return mapServiceDbToDomain(service), nil
}

func (s *ServiceGateway) GetAllServices() ([]models.Service, error) {

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
		mapped := mapServiceDbToDomain(serviceDb)
		mappedDbs = append(mappedDbs, mapped)
	}
	return mappedDbs, nil
}
