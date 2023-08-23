package admin

import (
	"encoding/json"
	"errors"
	"net/http"

	sg "github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/service"
	"github.com/rs/zerolog/log"
)

var (
	ErrBadGateway = errors.New("something went wrong in the server")
)

type ServiceManager struct {
	serviceGateway sg.ServiceGateway
}

func NewServiceManager(serviceGateway sg.ServiceGateway) *ServiceManager {
	return &ServiceManager{
		serviceGateway: serviceGateway,
	}
}

func (sm *ServiceManager) GetAllServicesHandler(w http.ResponseWriter, req *http.Request) {
	services, err := sm.serviceGateway.GetAllServices()
	if err != nil {
		log.Info().Msg(ErrBadGateway.Error())
		http.Error(w, ErrBadGateway.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(services)
}
