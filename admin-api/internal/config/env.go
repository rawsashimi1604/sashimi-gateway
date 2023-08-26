package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type EnvVars struct {
	POSTGRES_URL string
}

func LoadEnv() EnvVars {
	log.Info().Msg("Loading environment variables...")
	godotenv.Load()
	postgresUrl := os.Getenv("POSTGRES_URL")
	if postgresUrl == "" {
		log.Info().Msg("POSTGRES_URL: null")
	}

	return EnvVars{
		POSTGRES_URL: postgresUrl,
	}
}
