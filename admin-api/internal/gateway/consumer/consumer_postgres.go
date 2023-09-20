package consumer

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/utils"
	"github.com/rs/zerolog/log"
)

func NewPostgresConsumerGateway(conn *pgxpool.Pool) *PostgresConsumerGateway {
	return &PostgresConsumerGateway{Conn: conn}
}

func (cg *PostgresConsumerGateway) AddConsumerServices(cmer models.Consumer, servicesId []int) error {
	// Bulk insert using COPY command from Buffer
	var byteBuffer bytes.Buffer
	for _, serviceId := range servicesId {
		fmt.Fprintf(
			&byteBuffer,
			"%s\t%d\n",
			cmer.Id,
			serviceId,
		)
	}

	copyBuffer := utils.NewCopyBuffer(&byteBuffer)
	_, err := cg.Conn.CopyFrom(context.Background(),
		pgx.Identifier{"consumers_has_services"},
		[]string{"consumer_id", "service_id"},
		copyBuffer,
	)

	if err != nil {
		log.Info().Msg("something went wrong when adding bulk requests")
		return err
	}

	return nil
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

func (cg *PostgresConsumerGateway) GetConsumerById(id uuid.UUID) (models.Consumer, error) {

	query := `
	SELECT c.id, c.username, c.created_at, c.updated_at
	FROM consumer c
	WHERE c.id=$1
	`

	rows, err := cg.Conn.Query(context.Background(), query, id.String())
	if err != nil {
		return models.Consumer{}, err
	}
	defer rows.Close()

	consumerExists := false
	var consumer models.Consumer
	for rows.Next() {
		consumerExists = true
		var consDb Consumer_DB

		if err := rows.Scan(&consDb.Id, &consDb.Username, &consDb.CreatedAt, &consDb.UpdatedAt); err != nil {
			return models.Consumer{}, errors.New("error retrieving consumer")
		}

		consumer = MapConsumerDbToDomain(consDb)
	}

	if err := rows.Err(); err != nil {
		return models.Consumer{}, err
	}

	if !consumerExists {
		return models.Consumer{}, ErrConsumerNotFound
	}

	return consumer, nil
}
