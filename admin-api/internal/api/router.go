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

	conn := setupPostgresConn()
	reverseProxy := setupReverseProxy(conn)

	router := mux.NewRouter()
	router.Use(analytics.AnalyticsMiddleware)
	router.Use(reverseProxy.ReverseProxyMiddlware)

	log.Info().Msg("Admin Api Router created successfully.")
	return router
}

func setupPostgresConn() *pgxpool.Pool {
	conn, err := db.CreatePostgresConnection()
	if err != nil {
		log.Fatal().Msg("Unable to create postgres connection.")
	}
	return conn
}

func setupReverseProxy(conn *pgxpool.Pool) *rproxy.ReverseProxy {
	pgServiceGateway := service.NewPostgresServiceGateway(conn)
	return rproxy.NewReverseProxy(pgServiceGateway, http.DefaultTransport)

}
