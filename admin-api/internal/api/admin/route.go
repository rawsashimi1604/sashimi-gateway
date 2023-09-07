package admin

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	rg "github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/route"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/validator"
	"github.com/rs/zerolog/log"
)

var (
	ErrInvalidRouteBody = errors.New("invalid route body")
)

type RouteManager struct {
	routeGateway rg.RouteGateway
}

func NewRouteManager(routeGateway rg.RouteGateway) *RouteManager {
	return &RouteManager{
		routeGateway: routeGateway,
	}
}

func (rm *RouteManager) GetAllRoutesHandler(w http.ResponseWriter, req *http.Request) {
	routes, err := rm.routeGateway.GetAllRoutes()
	if err != nil {
		log.Info().Msg(ErrBadServer.Error())
		http.Error(w, ErrBadServer.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"count":  len(routes),
		"routes": routes,
	})
}

func (rm *RouteManager) RegisterRouteHandler(w http.ResponseWriter, req *http.Request) {

	type RegisterRouteRequest struct {
		ServiceId   int    `json:"serviceId" validate:"required"`
		Path        string `json:"path" validate:"required"`
		Description string `json:"description" validate:"required"`
		Method      string `json:"method" validate:"required"`
	}

	var body = RegisterRouteRequest{}

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Info().Msg(err.Error())
		log.Info().Msg(ErrInvalidRouteBody.Error())
		http.Error(w, ErrInvalidRouteBody.Error(), http.StatusBadRequest)
		return
	}

	validator := validator.NewValidator()
	err = validator.ValidateStruct(&body)
	if err != nil {
		log.Info().Msg(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: add Method gateway search for id.
	// TODO: add Service gateway to search for id.

	route := models.Route{
		Path:        body.Path,
		ServiceId:   body.ServiceId,
		Description: body.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Method:      body.Method,
	}

	registeredRoute, err := rm.routeGateway.RegisterRoute(route)
	if err != nil {
		log.Info().Msg(ErrBadServer.Error())
		http.Error(w, ErrBadServer.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(registeredRoute)
}
