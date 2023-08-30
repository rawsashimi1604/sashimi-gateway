package headers

import (
	"net/http"

	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/config"
)

func SetAdminHeadersMiddleware(next http.Handler) http.Handler {
	env := config.LoadEnv()
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", env.MANAGER_URL)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, req)
	})
}
