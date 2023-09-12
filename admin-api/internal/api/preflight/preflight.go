package preflight

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func HandlePreflightReqMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "OPTIONS" {
			log.Info().Msg("Preflight request received: OK")
			w.WriteHeader(http.StatusOK)
			return
		}
	})
}
