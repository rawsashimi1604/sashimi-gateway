package admin

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	sg "github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/service"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/utils"
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
		"services": services,
	})
}

func (sm *ServiceManager) RegisterServiceHandler(w http.ResponseWriter, req *http.Request) {

	type RegisterServiceRequest struct {
		Name        string `json:"name" validate:"required"`
		TargetUrl   string `json:"targetUrl" validate:"required"`
		Path        string `json:"path" validate:"required"`
		Description string `json:"description" validate:"required"`
	}

	var body = RegisterServiceRequest{}

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Info().Msg(ErrInvalidServiceBody.Error())
		http.Error(w, ErrInvalidServiceBody.Error(), http.StatusBadRequest)
		return
	}

	log.Info().Msg("service to be added: " + utils.JSONStringify(body))

	validator := validator.NewValidator()
	err = validator.ValidateStruct(&body)
	if err != nil {
		log.Info().Msg(ErrInvalidServiceBody.Error())
		http.Error(w, ErrInvalidServiceBody.Error(), http.StatusBadRequest)
		return
	}

	service := models.Service{
		Name:        body.Name,
		TargetUrl:   body.TargetUrl,
		Path:        body.Path,
		Description: body.Description,
		Routes:      make([]models.Route, 0),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// TODO: add check when service path already exists or is a RESERVED namespace.
	service, err = sm.serviceGateway.RegisterService(service)
	if err != nil {
		log.Info().Msg("something went wrong when creating service")
		http.Error(w, "something went wrong when creating service", http.StatusBadGateway)
		return
	}

	log.Info().Msg("service: " + utils.JSONStringify(service))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"service": service,
	})
}
