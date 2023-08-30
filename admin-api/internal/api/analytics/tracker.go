package analytics

import (
	"net/http"
	"sync"
	"time"

	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/utils"
	"github.com/rs/zerolog/log"
)

type AnalyticsTracker struct {
	requests []models.ApiRequest
	mutex    sync.Mutex
}

func NewAnalyticsTracker() *AnalyticsTracker {
	return &AnalyticsTracker{
		requests: make([]models.ApiRequest, 0),
		mutex:    sync.Mutex{},
	}
}

func (rt *AnalyticsTracker) Add(request models.ApiRequest) {
	rt.mutex.Lock()
	defer rt.mutex.Unlock()
	rt.requests = append(rt.requests, request)
}

func (rt *AnalyticsTracker) GetAndReset() []models.ApiRequest {
	rt.mutex.Lock()
	defer rt.mutex.Unlock()

	currentRequests := rt.requests
	rt.requests = make([]models.ApiRequest, 0)
	return currentRequests
}

func CaptureRequestMiddleware(rt *AnalyticsTracker) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestData := models.ApiRequest{
				Path:   r.URL.Path,
				Method: r.Method,
				Time:   time.Now(),
			}
			// Store the request data safely using mutex locks, serve http
			rt.Add(requestData)
			log.Info().Msg("rt slice: " + utils.JSONStringify(rt.requests))
			next.ServeHTTP(w, r)
		})
	}
}
