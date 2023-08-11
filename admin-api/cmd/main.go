package main

import (
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/logger"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/rproxy"
	"github.com/rs/zerolog/log"
)

func main() {
	// Set up zerolog configs.
	logger.SetupLogger()
	api.NewRouter()
	rproxy.ReverseProxy()

	log.Info().Msg("Hello world from admin api.")
}
