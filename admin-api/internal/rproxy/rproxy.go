package rproxy

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type Response struct {
	Code       int         `json:"code"`
	ReceivedAt time.Time   `json:"receivedAt"`
	Data       interface{} `json:"data"`
}

func ForwardRequest(w http.ResponseWriter, req *http.Request) {
	log.Info().Msg("Reverse proxy received request: " + req.Host)

	// Forward the request
	log.Info().Msg(time.Now().String())
	log.Info().Msg(req.Host)
	log.Info().Msg(req.URL.Host)
	log.Info().Msg(req.URL.Scheme)
	log.Info().Msg(req.RequestURI)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Code:       http.StatusOK,
		ReceivedAt: time.Now(),
		Data:       "Hello world",
	})
}
