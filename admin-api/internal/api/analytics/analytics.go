package analytics

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

// TODO: Add one more middleware to track websockets to manager
func AnalyticsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		log.Info().Msg("Reverse proxy received request: " + req.Host + " for path: " + req.URL.Path)
		next.ServeHTTP(w, req)
		log.Info().Msg("client ip: " + req.RemoteAddr)
		end := time.Now()
		duration := end.Sub(start).Milliseconds()
		log.Info().Msgf("duration: %vms", duration)
	})
}
