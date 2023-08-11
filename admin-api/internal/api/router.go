package api

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func NewRouter() *mux.Router {
	log.Info().Msg("Creating Mux Router.")

	router := mux.NewRouter()

	log.Info().Msg("Mux Router created successfully.")

	return router
}
