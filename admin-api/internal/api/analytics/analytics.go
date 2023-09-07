package analytics

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

// TODO: Add one more middleware to track websockets to manager
func AnalyticsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Info().Msg("Reverse proxy received request: " + req.Host + " for path: " + req.URL.Path)
		next.ServeHTTP(w, req)
	})
}
