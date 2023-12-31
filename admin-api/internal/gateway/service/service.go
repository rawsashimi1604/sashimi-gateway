package service

import (
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
)

type ServiceGateway interface {
	GetServiceByPath(path string) (models.Service, error)
	GetServiceByTargetUrl(targetUrl string) (models.Service, error)
	GetServiceById(id int) (models.Service, error)
	GetAllServices() ([]models.Service, error)
	GetAllServicesMatchingConsumer(consumer models.Consumer) ([]models.Service, error)
	RegisterService(service models.Service) (models.Service, error)
	UpdateServicesHealth(serviceHealthMap map[int]string) error
}

type PostgresServiceGateway struct {
	Conn *pgxpool.Pool
}

type Service_DB struct {
	Id                 int       `json:"id"`
	Name               string    `json:"name"`
	TargetUrl          string    `json:"targetUrl"`
	Path               string    `json:"path"`
	Description        string    `json:"description"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	HealthCheckEnabled bool      `json:"healthCheckEnabled"`
	Health             string    `json:"health"`
}

var (
	ErrServiceNotFound = errors.New("service not found in the database")
)

func MapServiceDbToDomain(sdb Service_DB, r []models.Route) models.Service {
	return models.Service{
		Id:                 sdb.Id,
		Name:               sdb.Name,
		TargetUrl:          sdb.TargetUrl,
		Path:               sdb.Path,
		Description:        sdb.Description,
		CreatedAt:          sdb.CreatedAt,
		UpdatedAt:          sdb.UpdatedAt,
		HealthCheckEnabled: sdb.HealthCheckEnabled,
		Health:             sdb.Health,
		Routes:             r,
	}
}
