package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type EnvVars struct {
	POSTGRES_URL             string
	MANAGER_URL              string
	SASHIMI_GATEWAY_NAME     string
	SASHIMI_HOSTNAME         string
	SASHIMI_TAGLINE          string
	SASHIMI_REQUEST_INTERVAL int
	SASHIMI_LOCAL_PORT       string
	SASHIMI_ADMIN_USERNAME   string
	SASHIMI_ADMIN_PASSWORD   string
	SASHIMI_ADMIN_JWT_KEY    string
}

func LoadEnv() EnvVars {
	godotenv.Load()

	postgresUrl := os.Getenv("POSTGRES_URL")
	managerUrl := os.Getenv("MANAGER_URL")
	sashimiGatewayName := os.Getenv("SASHIMI_GATEWAY_NAME")
	sashimiHostname := os.Getenv("SASHIMI_HOSTNAME")
	sashimiTagline := os.Getenv("SASHIMI_TAGLINE")
	sashimiRequestInterval := os.Getenv("SASHIMI_REQUEST_INTERVAL")
	sashimiPort := os.Getenv("SASHIMI_LOCAL_PORT")
	sashimiAdminUsername := os.Getenv("SASHIMI_ADMIN_USERNAME")
	sashimiAdminPassword := os.Getenv("SASHIMI_ADMIN_PASSWORD")
	sashimiAdminJwtKey := os.Getenv("SASHIMI_ADMIN_JWT_KEY")

	// Validate the environment variables
	requestInterval, err := strconv.Atoi(sashimiRequestInterval)
	if err != nil {
		log.Panic().Msg("invalid env variable: SASHIMI_REQUEST_INTERVAL: " + sashimiRequestInterval)
	}

	return EnvVars{
		POSTGRES_URL:             postgresUrl,
		MANAGER_URL:              managerUrl,
		SASHIMI_GATEWAY_NAME:     sashimiGatewayName,
		SASHIMI_HOSTNAME:         sashimiHostname,
		SASHIMI_TAGLINE:          sashimiTagline,
		SASHIMI_REQUEST_INTERVAL: requestInterval,
		SASHIMI_LOCAL_PORT:       sashimiPort,
		SASHIMI_ADMIN_USERNAME:   sashimiAdminUsername,
		SASHIMI_ADMIN_PASSWORD:   sashimiAdminPassword,
		SASHIMI_ADMIN_JWT_KEY:    sashimiAdminJwtKey,
	}
}
