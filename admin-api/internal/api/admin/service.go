package admin

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
	json.NewEncoder(w).Encode(map[string]interface{}{
		"count":    len(services),
		"services": services,
	})
}

func (sm *ServiceManager) GetServiceById(w http.ResponseWriter, req *http.Request) {
	serviceId := mux.Vars(req)["id"]
	serviceIdConverted, err := strconv.Atoi(serviceId)
	if err != nil {
		log.Info().Msg(err.Error())
		http.Error(w, "invalid id passed in", http.StatusBadRequest)
		return
	}

	service, err := sm.serviceGateway.GetServiceById(serviceIdConverted)
	if err != nil {
		if err == sg.ErrServiceNotFound {
			log.Info().Msg(sg.ErrServiceNotFound.Error())
			http.Error(w, sg.ErrServiceNotFound.Error(), http.StatusNotFound)
			return
		}
		log.Info().Msg("unable to retrieve service by id")
		http.Error(w, "unable to retrieve service by id", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"service": service,
	})
}

func (sm *ServiceManager) RegisterServiceHandler(w http.ResponseWriter, req *http.Request) {

	type RegisterServiceRequest struct {
		Name               string `json:"name" validate:"required"`
		TargetUrl          string `json:"targetUrl" validate:"required"`
		Path               string `json:"path" validate:"required"`
		Description        string `json:"description" validate:"required"`
		HealthCheckEnabled bool   `json:"healthCheckEnabled" validate:"required"`
	}

	var body = RegisterServiceRequest{}

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Info().Msg(err.Error())
		log.Info().Msg(ErrInvalidServiceBody.Error())
		http.Error(w, ErrInvalidServiceBody.Error(), http.StatusBadRequest)
		return
	}

	validator := validator.NewValidator()
	err = validator.ValidateStruct(&body)
	if err != nil {
		log.Info().Msg(ErrInvalidServiceBody.Error())
		http.Error(w, ErrInvalidServiceBody.Error(), http.StatusBadRequest)
		return
	}

	var health string
	if body.HealthCheckEnabled {
		health = "startup"
	} else {
		health = "not_enabled"
	}

	service := models.Service{
		Name:               body.Name,
		TargetUrl:          body.TargetUrl,
		Path:               body.Path,
		Description:        body.Description,
		Routes:             make([]models.Route, 0),
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		HealthCheckEnabled: body.HealthCheckEnabled,
		Health:             health,
	}

	// TODO: add check when service path already exists or is a RESERVED namespace.
	service, err = sm.serviceGateway.RegisterService(service)
	if err != nil {
		log.Info().Msg("something went wrong when creating service")
		http.Error(w, "something went wrong when creating service", http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"service": service,
	})
}
