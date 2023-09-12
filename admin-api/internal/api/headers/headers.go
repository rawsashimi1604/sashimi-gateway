package headers

import (
	"net/http"

	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/config"
	"github.com/rs/zerolog/log"
)

func SetAdminHeadersMiddleware(next http.Handler) http.Handler {
	env := config.LoadEnv()
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", env.MANAGER_URL)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		// Handler preflight requests.
		if req.Method == "OPTIONS" {
			log.Info().Msg("Preflight request received: OK")
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, req)
	})
}
