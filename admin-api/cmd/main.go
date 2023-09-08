package main

import (
	"net/http"

	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/config"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/logger"
	"github.com/rs/zerolog/log"
)

func main() {

	// TODO: add middleware for analytics (ongoing)
	// TODO: add health check route
	// TODO: add authentication (JWT)
	// TODO: add caching for services and routes (REDIS)
	// TODO: add rate limiting
	// TODO: add admin api (ongoing)
	// TODO: add GUI dashboard (ongoing)
	// TODO: refactor some services into their own seperate microservice.
	// TODO: track how long each request took.

	logger.SetupLogger()
	router := api.NewRouter()
	env := config.LoadEnv()

	log.Info().Msg("starting the admin api.")
	log.Info().Msg("admin api now listening for requests.")

	if err := http.ListenAndServe(":"+env.SASHIMI_LOCAL_PORT, router); err != nil {
		log.Fatal().Msg("error when starting the server.")
	}

}
