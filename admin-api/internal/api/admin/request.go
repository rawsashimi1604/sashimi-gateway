package admin

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	rq "github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/request"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/validator"
	"github.com/rs/zerolog/log"
)

var (
	ErrTimespanInvalid   = errors.New("no unsigned integer: timespan was given in the query")
	ErrDatapointsInvalid = errors.New("no unsigned integer: data points was given in the query")
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

	timespan := req.URL.Query().Get("timespan")
	dataPoints := req.URL.Query().Get("dataPoints")

	validator := validator.NewValidator()

	err := validator.ValidateSimple(timespan, "required,numeric")
	if err != nil {
		log.Info().Msg(ErrTimespanInvalid.Error())
		http.Error(w, ErrTimespanInvalid.Error(), http.StatusBadRequest)
		return
	}

	err = validator.ValidateSimple(dataPoints, "required,numeric")
	if err != nil {
		log.Info().Msg(ErrDatapointsInvalid.Error())
		http.Error(w, ErrDatapointsInvalid.Error(), http.StatusBadRequest)
		return
	}

	// Check for negative number inputs
	timeSpanConverted, err := strconv.ParseUint(timespan, 10, 32)
	if err != nil {
		log.Info().Msg(ErrTimespanInvalid.Error())
		http.Error(w, ErrTimespanInvalid.Error(), http.StatusBadRequest)
		return
	}

	dataPointsConverted, err := strconv.ParseUint(dataPoints, 10, 32)
	if err != nil {
		log.Info().Msg(ErrDatapointsInvalid.Error())
		http.Error(w, ErrDatapointsInvalid.Error(), http.StatusBadRequest)
		return
	}

	log.Info().Msg("timespan: " + timespan)
	log.Info().Msg("dataPoints: " + dataPoints)

	aggregatedRequests, err := reqm.requestGateway.GetAggregatedRequests(
		int(timeSpanConverted),
		int(dataPointsConverted),
	)

	if err != nil {
		log.Info().Msg(err.Error())
		http.Error(w, "error retrieving aggregated requests", http.StatusInternalServerError)
		return
	}

	dataPointsInt, _ := strconv.Atoi(dataPoints)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"dataPoints": dataPointsInt,
		"requests":   aggregatedRequests,
	})
}
