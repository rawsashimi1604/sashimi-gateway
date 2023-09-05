package admin

import (
	"encoding/json"
	"net/http"

	rq "github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/request"
	"github.com/rs/zerolog/log"
)

type RequestManager struct {
	requestGateway rq.RequestGateway
}

func NewRequestManager(requestGateway rq.RequestGateway) *RequestManager {
	return &RequestManager{
		requestGateway: requestGateway,
	}
}

func (reqm *RequestManager) GetAllRequestsHandler(w http.ResponseWriter, req *http.Request) {
	requests, err := reqm.requestGateway.GetAllRequests()
	if err != nil {
		log.Info().Msg(err.Error())
		http.Error(w, "error retrieving requests", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"count":    len(requests),
		"requests": requests,
	})
}

func (reqm *RequestManager) GetAggregatedRequestData(w http.ResponseWriter, req *http.Request) {

	DEFAULT_REQUEST_COUNT := 6

	aggregatedRequests, err := reqm.requestGateway.GetAggregatedRequests(15, 10)
	if err != nil {
		log.Info().Msg(err.Error())
		http.Error(w, "error retrieving aggregated requests", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"count":    DEFAULT_REQUEST_COUNT,
		"requests": aggregatedRequests,
	})
}
