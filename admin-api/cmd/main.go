package main

import (
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/config"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/logger"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/rproxy"
	"github.com/rs/zerolog/log"
)

func main() {
	env := config.LoadEnv()
	log.Info().Msg("Postgres URL loaded: " + env.POSTGRES_URL)
	logger.SetupLogger()
	api.NewRouter()
	rproxy.ReverseProxy()
	log.Info().Msg("Hello world from admin api.")

}
