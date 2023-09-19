package api

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	admin "github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api/admin"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api/analytics"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api/headers"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api/health"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api/rproxy"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/config"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/db"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/consumer"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/jwt_credentials"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/request"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/route"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/service"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/jobs"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/websocket"
	"github.com/rs/zerolog/log"
)

func NewRouter() *mux.Router {
	log.Info().Msg("Creating Admin Api Router.")

	// Load environment variables
	env := config.LoadEnv()

	// Setup dependencies
	conn := setupPostgresConn()

	// Create websocket server
	ws := websocket.NewWebSocketServer()

	// Load initial gateway information object
	gatewayConfig := admin.LoadInitialGatewayInfo(env)

	pgServiceGateway := service.NewPostgresServiceGateway(conn)
	pgRouteGateway := route.NewPostgresRouteGateway(conn)
	pgRequestGateway := request.NewPostgresRequestGateway(conn)
	pgConsumerGateway := consumer.NewPostgresConsumerGateway(conn)
	pgJwtCredentialsGateway := jwt_credentials.NewPostgresJWTCredentialsGateway(conn)

	// Gateway pattern (persistence, db data)
	gatewayManager := admin.NewGatewayManager(gatewayConfig)
	serviceManager := admin.NewServiceManager(pgServiceGateway)
	routeManager := admin.NewRouteManager(pgRouteGateway)
	requestManager := admin.NewRequestManager(pgRequestGateway)
	consumerManager := admin.NewConsumerManager(pgConsumerGateway)
	jwtCredentialsManager := admin.NewJwtCredentialsManager(pgJwtCredentialsGateway)
	adminAuthManager := admin.NewAdminAuthManager([]byte(env.SASHIMI_ADMIN_JWT_KEY))

	// Other services
	analyticsTracker := analytics.NewAnalyticsTracker(pgRequestGateway)
	healthChecker := health.NewHealthChecker(pgServiceGateway)
	reverseProxy := rproxy.NewReverseProxy(pgServiceGateway, analyticsTracker, http.DefaultTransport)

	// Cron job to periodically add requests to the database.
	requestCronJob := jobs.NewRequestCronJob(analyticsTracker, time.Duration(env.SASHIMI_REQUEST_INTERVAL)*time.Second, ws)
	healthCheckCronJob := jobs.NewHealthCheckCronJob(healthChecker, time.Duration(5*time.Second))
	requestCronJob.Start()
	healthCheckCronJob.Start()

	router := mux.NewRouter()

	// These route wont go through the reverse proxy middlewares
	adminRouter := router.PathPrefix("/api/admin").Subrouter()
	adminRouter.HandleFunc("/ws", ws.HandleClient)
	// Set CORS policy for admin Router
	adminRouter.Use(headers.SetAdminHeadersMiddleware)
	adminRouter.HandleFunc("/metadata", gatewayManager.GetGatewayMetadata).Methods("GET", "OPTIONS")
	adminRouter.HandleFunc("/login", adminAuthManager.Login).Methods("POST", "OPTIONS")
	adminRouter.HandleFunc("/auth/private-jwt", adminAuthManager.GetPrivateJwt).Methods("GET", "OPTIONS")
	adminRouter.HandleFunc("/auth/credentials", jwtCredentialsManager.ListCredentialsHandler).Methods("GET", "OPTIONS")
	adminRouter.HandleFunc("/auth/credentials/{id}", jwtCredentialsManager.GetAllCredentialsByConsumerId).Methods("GET", "OPTIONS")
	adminRouter.HandleFunc("/service/{id:[0-9]+}", serviceManager.GetServiceById).Methods("GET", "OPTIONS")
	adminRouter.HandleFunc("/service/all", serviceManager.GetAllServicesHandler).Methods("GET", "OPTIONS")
	adminRouter.HandleFunc("/service", serviceManager.RegisterServiceHandler).Methods("POST", "OPTIONS")
	adminRouter.HandleFunc("/route/all", routeManager.GetAllRoutesHandler).Methods("GET", "OPTIONS")
	adminRouter.HandleFunc("/route", routeManager.RegisterRouteHandler).Methods("POST", "OPTIONS")
	adminRouter.HandleFunc("/request/all", requestManager.GetAllRequestsHandler).Methods("GET", "OPTIONS")
	adminRouter.HandleFunc("/request/aggregate", requestManager.GetAggregatedRequestData).Methods("GET", "OPTIONS")
	adminRouter.HandleFunc("/consumer", consumerManager.RegisterConsumerHandler).Methods("POST", "OPTIONS")
	adminRouter.HandleFunc("/consumer", consumerManager.ListConsumers).Methods("GET", "OPTIONS")

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
