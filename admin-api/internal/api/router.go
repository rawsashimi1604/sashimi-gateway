package api

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func NewRouter() *mux.Router {
	log.Info().Msg("Creating Admin Api Router.")
	router := mux.NewRouter()
	log.Info().Msg("Admin Api Router created successfully.")
	return router
}
