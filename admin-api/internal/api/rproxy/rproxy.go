package rproxy

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	sg "github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/service"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/utils"

	"github.com/rs/zerolog/log"
)

type ReverseProxy struct {
	serviceGateway sg.ServiceGateway
	transport      http.RoundTripper
}

func NewReverseProxy(serviceGateway sg.ServiceGateway, httpTransport http.RoundTripper) *ReverseProxy {
	return &ReverseProxy{
		serviceGateway: serviceGateway,
		transport:      httpTransport,
	}
}

func (rps *ReverseProxy) ReverseProxyMiddlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		// validate validatedService
		validatedService, err := rps.matchService(req.URL.Path)
		if err != nil {
			if err == sg.ErrServiceNotFound {
				log.Info().Msg(fmt.Sprintf("service with path: %v not found.", req.URL.Path))
				http.Error(w, sg.ErrServiceNotFound.Error(), http.StatusNotFound)
				return
			}
			log.Info().Msg("service unable to be validated")
			http.Error(w, "service unable to be validated", http.StatusBadGateway)
			return
		}
		// log.Info().Msg("service: " + utils.JSONStringify(validatedService))

		// validate route
		validatedRoute, _, err := rps.matchRoute(validatedService, rps.parseRoutePath(req.URL.Path))
		if err != nil {
			log.Info().Msg("unable to find route")
			http.Error(w, "unable to find route", http.StatusNotFound)
			return
		}
		// log.Info().Msg("route: " + utils.JSONStringify(validatedRoute))

		// create origin url
		origin, err := url.Parse(validatedService.TargetUrl + validatedRoute.Path)
		if err != nil {
			log.Info().Msg("unable to parse upstream service and route url")
			http.Error(w, "unable to parse upstream service and route url", http.StatusBadRequest)
			return
		}
		log.Info().Msg("origin url: " + validatedService.TargetUrl + req.URL.Path)

		rps.prepareAndServeHttp(w, origin, req)
		next.ServeHTTP(w, req)
	})
}

func (rps *ReverseProxy) prepareAndServeHttp(w http.ResponseWriter, origin *url.URL, req *http.Request) {
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
	proxy.ModifyResponse = func(resp *http.Response) error {
		// Read the body data (and handle any errors)
		originalBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		// Close and replace the resp.Body, after reading the stream, stream will be at end, you must replace the data.
		resp.Body.Close()
		resp.Body = io.NopCloser(bytes.NewBuffer(originalBody))

		// Now you can inspect (or modify) the original body data
		log.Info().Msg("Origin Response Body:" + string(originalBody))

		// Return nil to indicate success
		return nil
	}
	proxy.ServeHTTP(w, req)
}

func (rps *ReverseProxy) parseServicePath(path string) string {
	urlSeperatedStrings := strings.Split(path, "/")
	return urlSeperatedStrings[1]
}

func (rps *ReverseProxy) parseRoutePath(path string) string {
	urlSeperatedStrings := strings.Split(path, "/")
	pathUrl := strings.Join(urlSeperatedStrings[2:], "/")
	return "/" + pathUrl
}

func (rps *ReverseProxy) matchService(path string) (models.Service, error) {
	service, err := rps.serviceGateway.GetServiceByPath(rps.parseServicePath(path))
	log.Info().Msg(utils.JSONStringify(service))
	if err != nil {
		if err == sg.ErrServiceNotFound {
			return models.Service{}, sg.ErrServiceNotFound
		}
		return models.Service{}, errors.New("something went wrong")
	}
	return service, nil
}

func (rps *ReverseProxy) matchRoute(service models.Service, requestPath string) (models.Route, map[string]string, error) {
	for _, route := range service.Routes {
		if isMatch, pathParams := rps.isRouteMatch(route.Path, requestPath); isMatch {
			return route, pathParams, nil
		}
	}
	return models.Route{}, nil, errors.New("unable to match route from service object")
}

func (rps *ReverseProxy) isRouteMatch(routePath string, requestPath string) (bool, map[string]string) {
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
