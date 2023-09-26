package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func NewRouter() *mux.Router {
	log.Info().Msg("Creating Admin Api Router.")

	router := mux.NewRouter()

	// Define any middleware
	// <add middleware here>

	// Define empty handler to catch all requests.
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	log.Info().Msg("Rev Proxy router created successfully.")
	return router
}
