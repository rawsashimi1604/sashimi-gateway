package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type EnvVars struct {
	POSTGRES_URL string
	MANAGER_URL  string
}

func LoadEnv() EnvVars {
	log.Info().Msg("Loading environment variables...")
	godotenv.Load()

	postgresUrl := os.Getenv("POSTGRES_URL")
	managerUrl := os.Getenv("MANAGER_URL")

	return EnvVars{
		POSTGRES_URL: postgresUrl,
		MANAGER_URL:  managerUrl,
	}
}
