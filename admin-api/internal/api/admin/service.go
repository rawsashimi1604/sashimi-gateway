package admin

import (
	"encoding/json"
	"errors"
	"net/http"

	sg "github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/service"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/validator"
	"github.com/rs/zerolog/log"
)

var (
	ErrBadServer          = errors.New("something went wrong in the server")
	ErrInvalidServiceBody = errors.New("invalid service body")
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
		log.Info().Msg(ErrBadServer.Error())
		http.Error(w, ErrBadServer.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(services)
}

func (sm *ServiceManager) RegisterServiceHandler(w http.ResponseWriter, req *http.Request) {
	var newService = models.Service{}
	err := json.NewDecoder(req.Body).Decode(&newService)
	if err != nil {
		log.Info().Msg(err.Error())
		log.Info().Msg(ErrInvalidServiceBody.Error())
		http.Error(w, ErrInvalidServiceBody.Error(), http.StatusBadRequest)
		return
	}

	validator := validator.NewValidator()
	err = validator.ValidateStruct(newService)
	if err != nil {
		log.Info().Msg(ErrInvalidServiceBody.Error())
		http.Error(w, ErrInvalidServiceBody.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newService)
}
