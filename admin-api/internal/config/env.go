package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	POSTGRES_URL string
}

func LoadEnv() EnvVars {
	godotenv.Load()
	postgresUrl := os.Getenv("POSTGRES_URL")
	return EnvVars{
		POSTGRES_URL: postgresUrl,
	}
}
