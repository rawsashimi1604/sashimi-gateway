package jwt_credentials

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
)

type JWTCredentialsGateway interface {
	ListCredentials() ([]models.JWTCredentials, error)
	GetAllCredentialsByConsumer(consumerId uuid.UUID) ([]models.JWTCredentials, error)
	AddCredential(models.JWTCredentials) (models.JWTCredentials, error)
}

type PostgresJWTCredentialsGateway struct {
	Conn *pgxpool.Pool
}

type JWTCredentials_DB struct {
	Id        string
	Key       string
	Secret    string
	Name      string
	CreatedAt time.Time
}

func MapJWTCredsDBToDomain(jcdb JWTCredentials_DB, consumer models.Consumer) models.JWTCredentials {
	id, _ := uuid.Parse(jcdb.Id)

	return models.JWTCredentials{
		Id:        id,
		Key:       jcdb.Key,
		Secret:    jcdb.Secret,
		Name:      jcdb.Name,
		Consumer:  consumer,
		CreatedAt: jcdb.CreatedAt,
	}
}
