package rproxy

import (
	"bytes"
	"errors"
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

// Routes:
// /salmon/products/:id -> match to
// TODO: create route matching algorithm (paths should be matched by ':' prefix... add it into the init.sql)

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

	pathUrl := rps.parseRoutePath(req.URL.Path)
	log.Info().Msg("path URL: " + pathUrl)

	service, err := rps.validateServiceExists(w, req.URL.Path)
	if err != nil {
		if err == gatewayService.ErrServiceNotFound {
			http.Error(w, gatewayService.ErrServiceNotFound.Error(), http.StatusNotFound)
			return
		}
		log.Info().Msg("service unable to be validated")
		http.Error(w, "service unable to be validated", http.StatusBadGateway)
		return
	}
	serviceURL, err := url.Parse(service.TargetUrl)
	if err != nil {
		log.Info().Msg("invalid url passed in.")
		http.Error(w, "invalid url passed in", http.StatusBadRequest)
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
		log.Info().Msg(err.Error())
		log.Info().Msg("something went wrong when forwarding the request")
		http.Error(w, "forward request error", http.StatusBadGateway)
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

func (rps *ReverseProxyService) parseServicePath(path string) string {
	urlSeperatedStrings := strings.Split(path, "/")
	return urlSeperatedStrings[1]
}

// join from index >= 2
func (rps *ReverseProxyService) parseRoutePath(path string) string {
	// Get from index >= 2
	// slice = ["", <service>, ....routes ]
	// example return : /products/1
	urlSeperatedStrings := strings.Split(path, "/")
	pathUrl := strings.Join(urlSeperatedStrings[2:], "/")
	return "/" + pathUrl
}

func (rps *ReverseProxyService) validateServiceExists(w http.ResponseWriter, path string) (models.Service, error) {
	service, err := rps.serviceGateway.GetServiceByPath(rps.parseServicePath(path))
	if err != nil {
		if err == gatewayService.ErrServiceNotFound {
			log.Info().Msg(fmt.Sprintf("service with path: %v not found.", path))
			return models.Service{}, gatewayService.ErrServiceNotFound
		}
		log.Info().Msg("Something went wrong")
		return models.Service{}, errors.New("something went wrong")
	}

	log.Info().Msg("service: " + utils.JSONStringify(service))
	return service, nil
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
