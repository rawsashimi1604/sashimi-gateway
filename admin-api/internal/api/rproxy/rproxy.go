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

	requestPathUrl := rps.parseRoutePath(req.URL.Path)

	service, err := rps.matchService(req.URL.Path)
	if err != nil {
		if err == gatewayService.ErrServiceNotFound {
			log.Info().Msg(fmt.Sprintf("service with path: %v not found.", req.URL.Path))
			http.Error(w, gatewayService.ErrServiceNotFound.Error(), http.StatusNotFound)
			return
		}
		log.Info().Msg("service unable to be validated")
		http.Error(w, "service unable to be validated", http.StatusBadGateway)
		return
	}

	validatedRoute, pathParams, err := rps.matchRoute(service, requestPathUrl)
	if err != nil {
		log.Info().Msg("unable to find route")
		http.Error(w, "unable to find route", http.StatusNotFound)
		return
	}

	log.Info().Msg("path URL: " + requestPathUrl)
	if req.URL.Query().Encode() != "" {
		log.Info().Msg("query params: " + req.URL.Query().Encode())
	} else {
		log.Info().Msg("query params: " + utils.JSONStringify(utils.EmptyStruct{}))
	}
	log.Info().Msg("path params: " + utils.JSONStringify(pathParams))
	log.Info().Msg("route validated: " + utils.JSONStringify(validatedRoute))

	serviceURL, err := url.Parse(service.TargetUrl)
	if err != nil {
		log.Info().Msg("invalid url passed in.")
		http.Error(w, "invalid url passed in", http.StatusBadRequest)
	}

	rps.modifyRequestHeaders(serviceURL, req, requestPathUrl)
	reqBodyBytes, err := utils.ReadHttpBody(req.Body)
	if err != nil {
		http.Error(w, "unable to read request body", http.StatusBadRequest)
		return
	}
	log.Info().Msg("request body: " + string(reqBodyBytes))

	// Reached end of stream when reading req.Body at the start, so set the req.Body again.
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

func (rps *ReverseProxyService) parseRoutePath(path string) string {
	// Get from index >= 2
	// slice = ["", <service>, ....routes ]
	// example return : /products/1
	urlSeperatedStrings := strings.Split(path, "/")
	pathUrl := strings.Join(urlSeperatedStrings[2:], "/")
	return "/" + pathUrl
}

func (rps *ReverseProxyService) modifyRequestHeaders(serviceURL *url.URL, req *http.Request, routePath string) {
	req.Host = serviceURL.Host
	req.URL.Host = serviceURL.Host
	req.URL.Scheme = serviceURL.Scheme
	req.URL.Path = routePath
	// We can't have this set when using http.DefaultClient
	req.RequestURI = ""
}

func (rps *ReverseProxyService) matchService(path string) (models.Service, error) {
	service, err := rps.serviceGateway.GetServiceByPath(rps.parseServicePath(path))
	if err != nil {
		if err == gatewayService.ErrServiceNotFound {
			return models.Service{}, gatewayService.ErrServiceNotFound
		}
		return models.Service{}, errors.New("something went wrong")
	}
	return service, nil
}

func (rps *ReverseProxyService) matchRoute(service models.Service, requestPath string) (models.Route, map[string]string, error) {
	for _, route := range service.Routes {
		if isMatch, pathParams := rps.isRouteMatch(route.Path, requestPath); isMatch {
			return route, pathParams, nil
		}
	}
	return models.Route{}, nil, errors.New("unable to match route from service object")
}

func (rps *ReverseProxyService) isRouteMatch(routePath string, requestPath string) (bool, map[string]string) {

	routeSegments := strings.Split(routePath, "/")
	requestSegments := strings.Split(requestPath, "/")

	if len(routeSegments) != len(requestSegments) {
		return false, nil
	}

	params := make(map[string]string)

	for i := range routeSegments {
		if strings.HasPrefix(routeSegments[i], ":") {
			paramKey := strings.TrimPrefix(routeSegments[i], ":")
			params[paramKey] = requestSegments[i]
		} else if routeSegments[i] != requestSegments[i] {
			return false, nil
		}
	}

	return true, params
}
