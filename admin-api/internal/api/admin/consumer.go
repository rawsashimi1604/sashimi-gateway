package admin

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	cg "github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/consumer"
	sv "github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/service"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/utils"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/validator"
	"github.com/rs/zerolog/log"
)

var (
	ErrInvalidConsumerBody = errors.New("invalid consumer body")
)

type ConsumerManager struct {
	consumerGateway cg.ConsumerGateway
	serviceGateway  sv.ServiceGateway
}

func NewConsumerManager(consumerGateway cg.ConsumerGateway, serviceGateway sv.ServiceGateway) *ConsumerManager {
	return &ConsumerManager{
		consumerGateway: consumerGateway,
		serviceGateway:  serviceGateway,
	}
}

func (cm *ConsumerManager) ListConsumers(w http.ResponseWriter, req *http.Request) {

	consumers, err := cm.consumerGateway.ListConsumers()
	if err != nil {
		log.Info().Msg(ErrBadServer.Error())
		http.Error(w, ErrBadServer.Error(), http.StatusInternalServerError)
		return
	}
	log.Info().Msg(utils.JSONStringify(consumers))

	services, err := cm.serviceGateway.GetAllServices()
	if err != nil {
		log.Info().Msg(ErrBadServer.Error())
		http.Error(w, ErrBadServer.Error(), http.StatusInternalServerError)
		return
	}

	resultDto := make([]map[string]interface{}, 0)

	log.Info().Msg(utils.JSONStringify(services))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"count":     len(resultDto),
		"consumers": resultDto,
	})
}

func (cm *ConsumerManager) RegisterConsumerHandler(w http.ResponseWriter, req *http.Request) {

	type RegisterConsumerRequest struct {
		Username string `json:"username" validate:"required"`
		Services []int  `json:"services" validate:"required,min=1"`
	}

	var body = RegisterConsumerRequest{}

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Info().Msg(ErrInvalidConsumerBody.Error())
		http.Error(w, ErrInvalidConsumerBody.Error(), http.StatusBadRequest)
		return
	}

	validator := validator.NewValidator()
	err = validator.ValidateStruct(&body)
	if err != nil {
		log.Info().Msg(ErrInvalidConsumerBody.Error())
		http.Error(w, ErrInvalidConsumerBody.Error(), http.StatusBadRequest)
		return
	}

	consumer := models.Consumer{
		Id:        uuid.New(),
		Username:  body.Username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	created, err := cm.consumerGateway.RegisterConsumer(consumer)
	if err != nil {
		log.Info().Msg(err.Error())
		log.Info().Msg("something went wrong when creating consumer")
		http.Error(w, "something went wrong when creating consumer", http.StatusInternalServerError)
		return
	}

	err = cm.consumerGateway.AddConsumerServices(created, body.Services)
	if err != nil {
		log.Info().Msg(err.Error())
		log.Info().Msg("somethng went wrong registering services for consumer")
		http.Error(w, "something went wrong registering services for customer", http.StatusInternalServerError)
		return
	}

	services, err := cm.serviceGateway.GetAllServicesMatchingConsumer(consumer)
	if err != nil {
		log.Info().Msg(err.Error())
		log.Info().Msg("somethng went wrong when retrieving services")
		http.Error(w, "something went wrong when retreiving services", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"consumer": map[string]interface{}{
			"id":        consumer.Id,
			"username":  consumer.Username,
			"createdAt": consumer.CreatedAt,
			"updatedAt": consumer.UpdatedAt,
			"services":  services,
		},
	})
}

func servicesContainsId(slice []int, id int) bool {
	for _, i := range slice {
		if i == id {
			return true
		}
	}
	return false
}
