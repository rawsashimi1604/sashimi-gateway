package consumer

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rs/zerolog/log"
)

func NewPostgresConsumerGateway(conn *pgxpool.Pool) *PostgresConsumerGateway {
	return &PostgresConsumerGateway{Conn: conn}
}

func (cg *PostgresConsumerGateway) RegisterConsumer(consumer models.Consumer) (models.Consumer, error) {
	query := `
	WITH inserted AS (
		INSERT INTO consumer
			(id, username, created_at, updated_at)
		VALUES 
			($1, $2, $3, $4)
		RETURNING *
	)
	SELECT id, username, created_at, updated_at
	FROM inserted
	`

	row := cg.Conn.QueryRow(
		context.Background(),
		query,
		consumer.Id,
		consumer.Username,
		consumer.CreatedAt,
		consumer.UpdatedAt,
	)

	createdConsumer := Consumer_DB{}

	if err := row.Scan(
		&createdConsumer.Id,
		&createdConsumer.Username,
		&createdConsumer.CreatedAt,
		&createdConsumer.UpdatedAt,
	); err != nil {
		log.Info().Msg(err.Error())
		return models.Consumer{}, nil
	}

	return MapConsumerDbToDomain(createdConsumer), nil
}
