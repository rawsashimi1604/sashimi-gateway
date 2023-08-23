package analytics

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func AnalyticsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Info().Msg("------------------")
		log.Info().Msg("Reverse proxy received request: " + req.Host + " for path: " + req.URL.Path)
		log.Info().Msg("Hello world from analytics middleware!")
		next.ServeHTTP(w, req)
	})
}
