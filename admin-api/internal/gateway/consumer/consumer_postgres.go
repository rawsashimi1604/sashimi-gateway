package consumer

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
)

func NewPostgresConsumerGateway(conn *pgxpool.Pool) *PostgresConsumerGateway {
	return &PostgresConsumerGateway{Conn: conn}
}

func (cg *PostgresConsumerGateway) RegisterConsumer(consumer models.Consumer) (models.Consumer, error) {
	return models.Consumer{}, nil
}
