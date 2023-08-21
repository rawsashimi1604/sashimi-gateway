package rproxy

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	gatewayService "github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/service"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/utils"

	"github.com/rs/zerolog/log"
)

type ReverseProxyService struct {
	serviceGateway gatewayService.ServiceGateway
}

func NewReverseProxyService(serviceGateway gatewayService.ServiceGateway) *ReverseProxyService {
	return &ReverseProxyService{
		serviceGateway: serviceGateway,
	}
}

func (rps *ReverseProxyService) ForwardRequest(w http.ResponseWriter, req *http.Request) {
	log.Info().Msg("------------------")
	log.Info().Msg("Reverse proxy received request: " + req.Host + " for path: " + req.URL.Path)

	service := rps.validateServiceExists(w, req.URL.Path)
	serviceURL, err := url.Parse(service.TargetUrl)
	if err != nil {
		log.Fatal().Msg("invalid url passed in.")
	}

	rps.modifyRequestHeaders(serviceURL, req)
	reqBodyBytes, err := utils.ReadHttpBody(req.Body)
	if err != nil {
		http.Error(w, "unable to read request body", http.StatusBadRequest)
		return
	}
	log.Info().Msg("request body: " + string(reqBodyBytes))

	// Send Http Request to the service
	// When you read from this stream (i.e., the request body), you are essentially consuming bytes from the beginning to the point you've read up to. Once you've read a byte, it's gone from the stream – you can't go back and read it again without some sort of intervention. Thats why we reset the reqBodyBytes, as the request has already been read from above.
	req.Body = io.NopCloser(bytes.NewReader(reqBodyBytes))
	serviceResponse, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Info().Msg(err.Error())
		log.Info().Msg("something went wrong when forwarding the request")
		return
	}

	respBodyBytes, err := utils.ReadHttpBody(serviceResponse.Body)
	if err != nil {
		http.Error(w, "Failed to read service response", http.StatusBadRequest)
		return
	}
	log.Info().Msg("response body: " + string(respBodyBytes))

	w.Header().Set("Content-Type", serviceResponse.Header.Get("Content-Type"))
	w.WriteHeader(http.StatusOK)
	w.Write(respBodyBytes)
}

func parseRequestPath(path string) string {
	urlSeperatedStrings := strings.Split(path, "/")
	return urlSeperatedStrings[1]
}

func (rps *ReverseProxyService) validateServiceExists(w http.ResponseWriter, path string) models.Service {
	service, err := rps.serviceGateway.GetServiceByPath(parseRequestPath(path))
	if err == gatewayService.ErrServiceNotFound {
		log.Info().Msg(fmt.Sprintf("service with path: %v not found.", path))
		http.Error(w, "service not found", http.StatusNotFound)
		return models.Service{}
	}
	if err != nil {
		log.Info().Msg(err.Error())
		log.Info().Msg("Something went wrong")
		http.Error(w, "something went wrong", http.StatusBadGateway)
		return models.Service{}
	}

	log.Info().Msg("service: " + utils.JSONStringify(service))
	return service
}

func (rps *ReverseProxyService) modifyRequestHeaders(serviceURL *url.URL, req *http.Request) {
	req.Host = serviceURL.Host
	req.URL.Host = serviceURL.Host
	req.URL.Scheme = serviceURL.Scheme
	// TODO: get the path to the server
	req.URL.Path = ""
	// We can't have this set when using http.DefaultClient
	req.RequestURI = ""
}
