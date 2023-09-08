package api

import (
	"net/http"
	"time"

	socketio "github.com/googollee/go-socket.io"
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
	"github.com/rs/zerolog/log"
)

func NewRouter() *mux.Router {
	log.Info().Msg("Creating Admin Api Router.")

	// Load environment variables
	env := config.LoadEnv()

	// Setup dependencies
	conn := setupPostgresConn()

	// Setup Websocket server
	websocketServer := setupWebSocketServer()

	// Load initial gateway information object
	gatewayConfig := admin.LoadInitialGatewayInfo(env)

	pgServiceGateway := service.NewPostgresServiceGateway(conn)
	pgRouteGateway := route.NewPostgresRouteGateway(conn)
	pgRequestGateway := request.NewPostgresRequestGateway(conn)

	// Gateway pattern (persistence, db data)
	gatewayManager := admin.NewGatewayManager(gatewayConfig)
	serviceManager := admin.NewServiceManager(pgServiceGateway)
	routeManager := admin.NewRouteManager(pgRouteGateway)
	requestManager := admin.NewRequestManager(pgRequestGateway)

	// Other services
	analyticsTracker := analytics.NewAnalyticsTracker(pgRequestGateway)
	reverseProxy := rproxy.NewReverseProxy(pgServiceGateway, analyticsTracker, http.DefaultTransport)

	// Cron job to periodically add requests to the database.
	requestCronJob := jobs.NewRequestCronJob(analyticsTracker, time.Duration(env.SASHIMI_REQUEST_INTERVAL)*time.Second, websocketServer)
	requestCronJob.Start()

	router := mux.NewRouter()

	// Create context middlware to pass to following req/res lifecycle.
	// router.Use(middleware.CreateContextMiddlware)

	// These route wont go through the reverse proxy middlewares
	adminRouter := router.PathPrefix("/api/admin").Subrouter()
	// Set CORS policy for admin Router
	adminRouter.Use(headers.SetAdminHeadersMiddleware)
	adminRouter.Handle("/socket.io/", websocketServer)
	adminRouter.HandleFunc("/general", gatewayManager.GetGatewayInformationHandler).Methods("GET")
	adminRouter.HandleFunc("/service/{id:[0-9]+}", serviceManager.GetServiceById).Methods("GET")
	adminRouter.HandleFunc("/service/all", serviceManager.GetAllServicesHandler).Methods("GET")
	adminRouter.HandleFunc("/service", serviceManager.RegisterServiceHandler).Methods("POST")
	adminRouter.HandleFunc("/route/all", routeManager.GetAllRoutesHandler).Methods("GET")
	adminRouter.HandleFunc("/route", routeManager.RegisterRouteHandler).Methods("POST")
	adminRouter.HandleFunc("/request/all", requestManager.GetAllRequestsHandler).Methods("GET")
	adminRouter.HandleFunc("/request/aggregate", requestManager.GetAggregatedRequestData).Methods("GET")

	// Other requests will go through the rproxy subrouter.
	reverseProxyRouter := router.PathPrefix("/").Subrouter()
	reverseProxyRouter.Use(analytics.AnalyticsMiddleware)
	reverseProxyRouter.Use(reverseProxy.ReverseProxyMiddleware) // Add the data to responses... use context

	// Define empty handler to catch all requests.
	reverseProxyRouter.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	log.Info().Msg("Admin Api Router created successfully.")
	return router
}

func setupWebSocketServer() *socketio.Server {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		log.Info().Msg("Connected!")
		s.SetContext("")
		s.Emit("connect", "connected to websocket server")
		s.Join("bcast")
		return nil
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		// Add the Remove session id. Fixed the connection & mem leak
		s.Emit("disconnect", "disconnected to websocket server")
	})

	log.Info().Msg("Created websocket server successfully.")
	go server.Serve()
	return server
}

func setupPostgresConn() *pgxpool.Pool {
	conn, err := db.CreatePostgresConnection()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	return conn
}
