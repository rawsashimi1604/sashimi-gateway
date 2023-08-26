package admin

import (
	"encoding/json"
	"net/http"

	rg "github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/route"
	"github.com/rs/zerolog/log"
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
	log.Info().Msg("Inside the get all routes handler")
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
