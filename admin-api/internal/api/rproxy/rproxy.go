package rproxy

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	gatewayService "github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/service"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"

	"github.com/rs/zerolog/log"
)

type ReverseProxyService struct {
	serviceGateway gatewayService.ServiceGateway
	transport      http.RoundTripper
}

func NewReverseProxyService(serviceGateway gatewayService.ServiceGateway, httpTransport http.RoundTripper) *ReverseProxyService {
	return &ReverseProxyService{
		serviceGateway: serviceGateway,
		transport:      httpTransport,
	}
}

func (rps *ReverseProxyService) ForwardRequest(w http.ResponseWriter, req *http.Request) {
	log.Info().Msg("------------------")
	log.Info().Msg("Reverse proxy received request: " + req.Host + " for path: " + req.URL.Path)

	// validate service
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

	// validate route
	validatedRoute, _, err := rps.matchRoute(service, rps.parseRoutePath(req.URL.Path))
	if err != nil {
		log.Info().Msg("unable to find route")
		http.Error(w, "unable to find route", http.StatusNotFound)
		return
	}

	// create origin url
	origin, err := url.Parse(service.TargetUrl + validatedRoute.Path)
	if err != nil {
		log.Info().Msg("unable to parse upstream service and route url")
		http.Error(w, "unable to parse upstream service and route url", http.StatusBadRequest)
		return
	}

	rps.prepareAndServeHttp(w, origin, req)
}

func (rps *ReverseProxyService) prepareAndServeHttp(w http.ResponseWriter, origin *url.URL, req *http.Request) {
	proxy := httputil.NewSingleHostReverseProxy(origin)
	proxy.Transport = rps.transport
	proxy.Director = func(directorReq *http.Request) {
		directorReq.Header.Add("X-Forwarded-Host", req.Host)
		directorReq.Header.Add("X-Origin-Host", origin.Host)
		directorReq.URL.Scheme = origin.Scheme
		directorReq.URL.Host = origin.Host
		directorReq.URL.Path = rps.parseRoutePath(req.URL.Path)
	}
	proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {
		log.Info().Msgf("Error while proxying request: %v", err)
		http.Error(w, "Error while proxying request", http.StatusInternalServerError)
	}
	proxy.ServeHTTP(w, req)
}

func (rps *ReverseProxyService) parseServicePath(path string) string {
	urlSeperatedStrings := strings.Split(path, "/")
	return urlSeperatedStrings[1]
}

func (rps *ReverseProxyService) parseRoutePath(path string) string {
	urlSeperatedStrings := strings.Split(path, "/")
	pathUrl := strings.Join(urlSeperatedStrings[2:], "/")
	return "/" + pathUrl
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
