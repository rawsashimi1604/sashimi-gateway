package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api/analytics"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api/rproxy"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/db"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/service"
	"github.com/rs/zerolog/log"
)

func NewRouter() *mux.Router {
	log.Info().Msg("Creating Admin Api Router.")

	router := mux.NewRouter()
	setupMiddleware(router)
	conn := setupPostgresConn()
	rproxyService := setupReverseProxyService(conn)

	router.PathPrefix("/").HandlerFunc(rproxyService.ForwardRequest).Methods("GET", "PUT", "POST", "DELETE")
	http.Handle("/", router)

	log.Info().Msg("Admin Api Router created successfully.")
	return router
}

func setupMiddleware(router *mux.Router) {
	router.Use(analytics.AnalyticsMiddleware)
}

func setupPostgresConn() *pgxpool.Pool {
	conn, err := db.CreatePostgresConnection()
	if err != nil {
		log.Fatal().Msg("Unable to create postgres connection.")
	}
	return conn
}

func setupReverseProxyService(conn *pgxpool.Pool) *rproxy.ReverseProxyService {
	pgServiceGateway := service.NewPostgresServiceGateway(conn)
	return rproxy.NewReverseProxyService(pgServiceGateway, http.DefaultTransport)

}
