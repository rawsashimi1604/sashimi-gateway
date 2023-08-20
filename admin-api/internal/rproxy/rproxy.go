package rproxy

import (
	"io"
	"net/http"
	"net/url"
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
	log.Info().Msg(time.Now().String())

	// Define service server
	serviceURL, err := url.Parse("http://localhost:8081")
	if err != nil {
		log.Fatal().Msg("invalid url passed in.")
	}

	// We can't have this set when using http.DefaultClient
	req.Host = serviceURL.Host
	req.URL.Host = serviceURL.Host
	req.URL.Scheme = serviceURL.Scheme
	req.RequestURI = ""

	serviceResponse, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Info().Msg(err.Error())
		log.Info().Msg("something went wrong when forwarding the request")
		return
	}

	w.WriteHeader(http.StatusOK)
	io.Copy(w, serviceResponse.Body)

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(Response{
	// 	Code:       http.StatusOK,
	// 	ReceivedAt: time.Now(),
	// 	Data:       "Hello world",
	// })
}
