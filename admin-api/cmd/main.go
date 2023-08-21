package main

import (
	"net/http"

	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/logger"
	"github.com/rs/zerolog/log"
)

func main() {

	// TODO: add middleware for analytics

	logger.SetupLogger()
	router := api.NewRouter()

	log.Info().Msg("starting the admin api.")
	log.Info().Msg("admin api now listening for requests.")

	http.Handle("/", router)
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal().Msg("error when starting the server.")
	}

}
