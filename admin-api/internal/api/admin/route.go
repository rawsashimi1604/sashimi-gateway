package admin

import (
	"encoding/json"
	"errors"
	"net/http"

	rg "github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/route"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/utils"
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

	w.Header().Set("Content-Type", "applciation/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(routes)
}

func (rm *RouteManager) RegisterRouteHandler(w http.ResponseWriter, req *http.Request) {

	type RegisterRouteRequest struct {
		ServiceId   int    `json:"serviceId" validate:"required"`
		Path        string `json:"path" validate:"required"`
		Description string `json:"description" validate:"required"`
		MethodId    int    `json:"methodId" validate:"required"`
	}

	var body = RegisterRouteRequest{}

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Info().Msg(ErrInvalidRouteBody.Error())
		http.Error(w, ErrInvalidRouteBody.Error(), http.StatusBadRequest)
		return
	}
	log.Info().Msg("route to be added: " + utils.JSONStringify(body))

	validator := validator.NewValidator()
	err = validator.ValidateStruct(&body)
	if err != nil {
		log.Info().Msg(ErrInvalidRouteBody.Error())
		http.Error(w, ErrInvalidRouteBody.Error(), http.StatusBadRequest)
		return
	}

	// TODO: add Method gateway to get all methods and search for id.

	// route := models.Route{
	// 	Path:        body.Path,
	// 	ServiceId:   body.ServiceId,
	// 	Description: body.Description,
	// 	CreatedAt:   time.Now(),
	// 	UpdatedAt:   time.Now(),
	// 	Method:      models.ApiMethod{Id: 1, Method: "GET"},
	// }

}
