package middleware

import (
	"context"
	"net/http"
)

type ContextKey string

// Create the context and pass it to the other middlewares.
func CreateContextMiddlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := context.Background()
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
