package rproxy

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/rs/zerolog/log"
)

func ForwardRequest(w http.ResponseWriter, req *http.Request) {
	log.Info().Msg("Reverse proxy received request: " + req.Host)
	log.Info().Msg(time.Now().String())

	// Read request body
	reqBodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		log.Info().Msg("Failed to read request body")
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	// You might want to log the request body here
	log.Info().Msg("body: " + string(reqBodyBytes))

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

	reqBody, err := json.Marshal(serviceResponse.Body)
	if err != nil {
		log.Info().Msg("something went wrong when marshalling json.")
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(reqBody)
	// io.Copy(w, serviceResponse.Body)
}
