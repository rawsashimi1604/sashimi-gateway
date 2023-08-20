package rproxy

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/db"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/mapper"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/utils"

	"github.com/rs/zerolog/log"
)

// TODO: add tables and parse.. match the service psql table...
func ForwardRequest(w http.ResponseWriter, req *http.Request) {
	log.Info().Msg("Reverse proxy received request: " + req.Host)
	log.Info().Msg(time.Now().String())
	log.Info().Msg(req.URL.Path)
	log.Info().Msg(req.Method)

	// Create Gateway to access services
	conn, err := db.CreatePostgresConnection()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	serviceGateway := gateway.NewServiceGateway(conn)
	servicesDb, err := serviceGateway.GetAllServices()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	// map service_db to service domain
	var services []*models.Service
	for _, service := range servicesDb {
		services = append(
			services,
			mapper.MapServiceDbToDomain(&service),
		)
	}

	log.Info().Msg(utils.JSONStringify(services))

	// Define service server
	serviceURL, err := url.Parse("http://localhost:8082")
	if err != nil {
		log.Fatal().Msg("invalid url passed in.")
	}

	req.Host = serviceURL.Host
	req.URL.Host = serviceURL.Host
	req.URL.Scheme = serviceURL.Scheme
	// We can't have this set when using http.DefaultClient
	req.RequestURI = ""

	// Read request body
	reqBodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		log.Info().Msg("Failed to read request body")
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	log.Info().Msg(fmt.Sprintf("request body: %q", func() string {
		if len(reqBodyBytes) == 0 {
			return "No body found."
		}
		return string(reqBodyBytes)
	}()))

	// Send Http Request to the service
	serviceResponse, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Info().Msg(err.Error())
		log.Info().Msg("something went wrong when forwarding the request")
		return
	}

	// Read the response body
	respBodyBytes, err := io.ReadAll(serviceResponse.Body)
	if err != nil {
		log.Info().Msg("Failed to read service response body")
		http.Error(w, "Failed to read service response", http.StatusInternalServerError)
		return
	}
	log.Info().Msg("response body: " + string(respBodyBytes))

	// Send back response
	w.Header().Set("Content-Type", serviceResponse.Header.Get("Content-Type"))
	w.WriteHeader(http.StatusOK)
	w.Write(respBodyBytes)
}
