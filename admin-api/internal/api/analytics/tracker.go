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

func (at *AnalyticsTracker) Add(request models.ApiRequest) {
	at.mutex.Lock()
	defer at.mutex.Unlock()
	at.requests = append(at.requests, request)
}

func (at *AnalyticsTracker) GetAndReset() []models.ApiRequest {
	at.mutex.Lock()
	defer at.mutex.Unlock()

	currentRequests := at.requests
	at.requests = make([]models.ApiRequest, 0)
	return currentRequests
}

func (at *AnalyticsTracker) CaptureRequest(serviceId int, routeId int, req *http.Request) {
	requestData := models.ApiRequest{
		ServiceId: serviceId,
		RouteId:   routeId,
		Path:      req.URL.Path,
		Method:    req.Method,
		Time:      time.Now(),
	}
	// Store the request data safely using mutex locks
	at.Add(requestData)
	log.Info().Msg("rt slice: " + utils.JSONStringify(at.requests))
}
