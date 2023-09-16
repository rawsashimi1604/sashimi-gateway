package analytics

import (
	"context"
	"net/http"
	"time"

	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/middleware"
	"github.com/rs/zerolog/log"
)

// TODO: Add one more middleware to track websockets to manager
func AnalyticsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Start the timer for request tracking
		start := time.Now().UnixMilli()
		ctx := context.WithValue(req.Context(), middleware.ApiRequestDuration, start)
		log.Info().Msg("Reverse proxy received request: " + req.Host + " for path: " + req.URL.Path)
		log.Info().Msg("client ip: " + req.RemoteAddr)
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
