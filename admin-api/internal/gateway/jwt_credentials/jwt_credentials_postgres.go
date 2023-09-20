package jwt_credentials

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/consumer"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rs/zerolog/log"
)

func NewPostgresJWTCredentialsGateway(conn *pgxpool.Pool) *PostgresJWTCredentialsGateway {
	return &PostgresJWTCredentialsGateway{Conn: conn}
}

func (jcg *PostgresJWTCredentialsGateway) ListCredentials() ([]models.JWTCredentials, error) {
	query := `
		SELECT jc.id, jc.key, jc.secret, jc.name, jc.consumer_id, jc.created_at, c.username, c.created_at, c.updated_at
		FROM jwt_credentials jc
		LEFT JOIN consumer c
		ON c.id=jc.consumer_id
		ORDER BY c.id ASC
	`

	rows, err := jcg.Conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var credentials []models.JWTCredentials
	for rows.Next() {
		var jwtCredentials JWTCredentials_DB
		var consumerDb consumer.Consumer_DB

		if err := rows.Scan(
			&jwtCredentials.Id,
			&jwtCredentials.Key,
			&jwtCredentials.Secret,
			&jwtCredentials.Name,
			&jwtCredentials.CreatedAt,
			&consumerDb.Id,
			&consumerDb.Username,
			&consumerDb.CreatedAt,
			&consumerDb.UpdatedAt,
		); err != nil {
			return nil, errors.New("error retrieving jwt credential")
		}

		credentials = append(credentials, MapJWTCredsDBToDomain(
			jwtCredentials,
			consumer.MapConsumerDbToDomain(consumerDb),
		))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return credentials, nil
}

func (jcg *PostgresJWTCredentialsGateway) GetAllCredentialsByConsumer(consumerId uuid.UUID) ([]models.JWTCredentials, error) {

	converted := consumerId.String()

	query := `
		SELECT jc.id, jc.key, jc.secret, jc.name, jc.consumer_id, jc.created_at, c.username, c.created_at, c.updated_at
		FROM jwt_credentials jc
		LEFT JOIN consumer c
		ON c.id=jc.consumer_id
		WHERE c.id=$1
		ORDER BY c.id ASC
	`

	rows, err := jcg.Conn.Query(context.Background(), query, converted)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var credentials []models.JWTCredentials
	for rows.Next() {
		var jwtCredentials JWTCredentials_DB
		var consumerDb consumer.Consumer_DB

		if err := rows.Scan(
			&jwtCredentials.Id,
			&jwtCredentials.Key,
			&jwtCredentials.Secret,
			&jwtCredentials.Name,
			&jwtCredentials.CreatedAt,
			&consumerDb.Id,
			&consumerDb.Username,
			&consumerDb.CreatedAt,
			&consumerDb.UpdatedAt,
		); err != nil {
			return nil, errors.New("error retrieving jwt credential")
		}

		credentials = append(credentials, MapJWTCredsDBToDomain(
			jwtCredentials,
			consumer.MapConsumerDbToDomain(consumerDb),
		))
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return credentials, nil
}

func (jcg *PostgresJWTCredentialsGateway) AddCredential(credential models.JWTCredentials) (models.JWTCredentials, error) {

	query := `
	WITH ins AS (
		INSERT INTO jwt_credentials(id, key, secret, name, consumer_id, created_at) 
		VALUES
			($1, $2, $3, $4, $5, $6)
		RETURNING *
	)
	SELECT 
		ins.id, ins.key, ins.secret, ins.name, ins.consumer_id, ins.created_at,
		c.id, c.username, c.created_at AS consumer_created_at, c.updated_at AS consumer_updated_at
	FROM ins
	JOIN consumer c ON ins.consumer_id = c.id;
	`

	row := jcg.Conn.QueryRow(
		context.Background(),
		query,
		credential.Id.String(),
		credential.Key,
		credential.Secret,
		credential.Name,
		credential.CreatedAt,
	)

	createdCredential := JWTCredentials_DB{}
	relatedConsumer := consumer.Consumer_DB{}

	if err := row.Scan(
		&createdCredential.Id,
		&createdCredential.Key,
		&createdCredential.Secret,
		&createdCredential.Name,
		&relatedConsumer.Id,
		&createdCredential.CreatedAt,
		&relatedConsumer.Username,
		&relatedConsumer.CreatedAt,
		&relatedConsumer.UpdatedAt,
	); err != nil {
		log.Info().Msg(err.Error())
		return models.JWTCredentials{}, err
	}

	// TODO: get the related consumer...
	return MapJWTCredsDBToDomain(
		createdCredential,
		consumer.MapConsumerDbToDomain(relatedConsumer),
	), nil
}
