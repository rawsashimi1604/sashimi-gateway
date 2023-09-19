package jwt_credentials

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/consumer"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
)

func NewPostgresJWTCredentialsGateway(conn *pgxpool.Pool) *PostgresJWTCredentialsGateway {
	return &PostgresJWTCredentialsGateway{Conn: conn}
}

func (jcg *PostgresJWTCredentialsGateway) ListCredentials() ([]models.JWTCredentials, error) {
	query := `
		SELECT jc.id, jc.key, jc.secret, jc.name, jc.consumer_id, c.username, c.created_at, c.updated_at
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
