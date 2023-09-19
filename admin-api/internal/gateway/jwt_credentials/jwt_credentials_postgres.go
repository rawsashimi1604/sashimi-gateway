package jwt_credentials

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresJWTCredentialsGateway(conn *pgxpool.Pool) *PostgresJWTCredentialsGateway {
	return &PostgresJWTCredentialsGateway{Conn: conn}
}
