package api

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	admin "github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api/admin"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api/analytics"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api/headers"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api/rproxy"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/config"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/db"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/request"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/route"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/service"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/jobs"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/middleware"
	"github.com/rs/zerolog/log"
)

func NewRouter() *mux.Router {
	log.Info().Msg("Creating Admin Api Router.")

	// Load environment variables
	env := config.LoadEnv()

	// Setup dependencies
	conn := setupPostgresConn()

	pgServiceGateway := service.NewPostgresServiceGateway(conn)
	pgRouteGateway := route.NewPostgresRouteGateway(conn)
	pgRequestGateway := request.NewPostgresRequestGateway(conn)

	// Gateway pattern (persistence, db data)
	gatewayManager := admin.NewGatewayManager()
	serviceManager := admin.NewServiceManager(pgServiceGateway)
	routeManager := admin.NewRouteManager(pgRouteGateway)

	// Other services
	analyticsTracker := analytics.NewAnalyticsTracker(pgRequestGateway)
	reverseProxy := rproxy.NewReverseProxy(pgServiceGateway, analyticsTracker, http.DefaultTransport)

	// Cron job to periodically add requests to the database.
	requestCronJob := jobs.NewRequestCronJob(analyticsTracker, time.Duration(env.SASHIMI_REQUEST_INTERVAL)*time.Second)
	requestCronJob.Start()

	router := mux.NewRouter()
	// Create context middlware to pass to following req/res lifecycle.
	router.Use(middleware.CreateContextMiddlware)

	// These route wont go through the reverse proxy middlewares
	adminRouter := router.PathPrefix("/api/admin").Subrouter()
	// Set CORS policy for admin Router
	adminRouter.Use(headers.SetAdminHeadersMiddleware)
	adminRouter.HandleFunc("/general", gatewayManager.GetGatewayInformationHandler).Methods("GET")
	adminRouter.HandleFunc("/service/all", serviceManager.GetAllServicesHandler).Methods("GET")
	adminRouter.HandleFunc("/service", serviceManager.RegisterServiceHandler).Methods("POST")
	adminRouter.HandleFunc("/route/all", routeManager.GetAllRoutesHandler).Methods("GET")
	adminRouter.HandleFunc("/route", routeManager.RegisterRouteHandler).Methods("POST")

	// Other requests will go through the rproxy subrouter.
	reverseProxyRouter := router.PathPrefix("/").Subrouter()
	reverseProxyRouter.Use(analytics.AnalyticsMiddleware)
	reverseProxyRouter.Use(reverseProxy.ReverseProxyMiddleware) // Add the data to responses... use context

	// Define empty handler to catch all requests.
	reverseProxyRouter.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	log.Info().Msg("Admin Api Router created successfully.")
	return router
}

func setupPostgresConn() *pgxpool.Pool {
	conn, err := db.CreatePostgresConnection()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	return conn
}
