package api

import (
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/rproxy"
	"github.com/rs/zerolog/log"
)

func NewRouter() *mux.Router {
	log.Info().Msg("Creating Admin Api Router.")
	router := mux.NewRouter()
	router.HandleFunc("/", rproxy.ForwardRequest).Methods("GET", "PUT", "POST", "DELETE")
	log.Info().Msg("Admin Api Router created successfully.")
	return router
}
