package analytics

import (
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/request"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/utils"
	"github.com/rs/zerolog/log"
)

type AnalyticsTracker struct {
	requests       []models.ApiRequest
	requestGateway request.RequestGateway
	mutex          sync.Mutex
}

func NewAnalyticsTracker(gateway request.RequestGateway) *AnalyticsTracker {
	return &AnalyticsTracker{
		requests:       make([]models.ApiRequest, 0),
		mutex:          sync.Mutex{},
		requestGateway: gateway,
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

func (at *AnalyticsTracker) CaptureRequest(service models.Service, route models.Route, req *http.Request, statusCode int, duration int64) {
	requestData := models.ApiRequest{
		Id:        uuid.New(),
		ServiceId: service.Id,
		RouteId:   route.Id,
		Path:      req.URL.Path,
		Method:    req.Method,
		Time:      time.Now(),
		Code:      statusCode,
		Duration:  duration,
	}
	// Store the request data safely using mutex locks
	at.Add(requestData)
}

func (at *AnalyticsTracker) StoreRequests() []models.ApiRequest {
	// Get the requests safely using mutexes lock and unlock mechanism
	requests := at.GetAndReset()
	_, err := at.requestGateway.AddBulkRequests(requests)
	log.Info().Msg("added the following requests to db: " + utils.JSONStringify(requests))
	if err != nil {
		log.Info().Msg(err.Error())
		log.Info().Msg("something went wrong when storing requests")
	}

	return requests
}
