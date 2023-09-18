package consumer

import (
	"context"
	"errors"

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

func (cg *PostgresConsumerGateway) ListConsumers() ([]models.Consumer, error) {
	query := `
		SELECT c.id, c.username, c.created_at, c.updated_at
		FROM consumer c
		ORDER BY c.id ASC
	`

	rows, err := cg.Conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var consumers []models.Consumer
	for rows.Next() {
		var consumer Consumer_DB

		if err := rows.Scan(&consumer.Id, &consumer.Username, &consumer.CreatedAt, &consumer.UpdatedAt); err != nil {
			return nil, errors.New("error retrieving consumer")
		}

		consumers = append(consumers, MapConsumerDbToDomain(consumer))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return consumers, nil
}
