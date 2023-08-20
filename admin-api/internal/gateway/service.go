package gateway

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
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

func NewServiceGateway(conn *pgxpool.Pool) *ServiceGateway {
	return &ServiceGateway{Conn: conn}
}

func (s *ServiceGateway) GetAllServices() ([]Service_DB, error) {

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
		log.Info().Msg(fmt.Sprintf("service with id: %v successfully retrieved.", &service.Id))
		services = append(services, service)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return services, nil
}
