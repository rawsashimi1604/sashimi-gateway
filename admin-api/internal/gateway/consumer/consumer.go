package consumer

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
)

type ConsumerGateway interface {
	RegisterConsumer(consumer models.Consumer) (models.Consumer, error)
	ListConsumers() ([]models.Consumer, error)
	AddConsumerServices(cmer models.Consumer, servicesId []int) error
	GetConsumerById(id uuid.UUID) (models.Consumer, error)
}

var (
	ErrConsumerNotFound = errors.New("consuemr not found in the database")
)

type PostgresConsumerGateway struct {
	Conn *pgxpool.Pool
}

type Consumer_DB struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func MapConsumerDbToDomain(cdb Consumer_DB) models.Consumer {
	id, _ := uuid.Parse(cdb.Id)

	return models.Consumer{
		Id:        id,
		Username:  cdb.Username,
		CreatedAt: cdb.CreatedAt,
		UpdatedAt: cdb.UpdatedAt,
	}
}
