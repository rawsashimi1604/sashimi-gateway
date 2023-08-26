package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/config"
	"github.com/rs/zerolog/log"
)

var (
	ErrCreatePGConnection = errors.New("something went wrong connecting to postgres, please check your connection string")
	ErrPingPG             = errors.New("something went wrong connecting to postgres, unable to ping database")
)

func CreatePostgresConnection() (*pgxpool.Pool, error) {
	env := config.LoadEnv()
	log.Info().Msg("Connecting to postgres...")
	conn, err := pgxpool.New(context.Background(), env.POSTGRES_URL)
	if err != nil {
		return nil, ErrCreatePGConnection
	}

	if err = conn.Ping(context.Background()); err != nil {
		return nil, ErrPingPG
	}
	log.Info().Msg("Postgres connected successfully.")
	return conn, nil
}
