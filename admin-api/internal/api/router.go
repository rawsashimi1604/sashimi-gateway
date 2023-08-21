package api

import (
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/db"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/rproxy"
	"github.com/rs/zerolog/log"
)

func NewRouter() *mux.Router {
	log.Info().Msg("Creating Admin Api Router.")

	router := mux.NewRouter()

	// Create postgres database connection
	conn, err := db.CreatePostgresConnection()
	if err != nil {
		log.Fatal().Msg("Unable to create postgres connection.")
	}

	// Inject gateway dependencies
	pgServiceGateway := gateway.NewPostgresServiceGateway(conn)
	rproxyService := rproxy.NewReverseProxyService(pgServiceGateway)

	router.PathPrefix("/").HandlerFunc(rproxyService.ForwardRequest).Methods("GET", "PUT", "POST", "DELETE")
	log.Info().Msg("Admin Api Router created successfully.")
	return router
}
