package main

import (
	"github.com/rawsashimi1604/sashimi-gateway/rev-proxy/internal/infra/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	logger.SetupLogger()
	log.Info().Msg("Hello world from logger.")
}
