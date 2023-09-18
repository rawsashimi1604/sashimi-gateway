package admin

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	cg "github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/consumer"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rs/zerolog/log"
)

var (
	ErrInvalidConsumerBody = errors.New("invalid consumer body")
)

type ConsumerManager struct {
	consumerGateway cg.ConsumerGateway
}

func NewConsumerManager(consumerGateway cg.ConsumerGateway) *ConsumerManager {
	return &ConsumerManager{
		consumerGateway: consumerGateway,
	}
}

func (cm *ConsumerManager) ListConsumers(w http.ResponseWriter, req *http.Request) {

	consumers, err := cm.consumerGateway.ListConsumers()
	if err != nil {
		log.Info().Msg(ErrBadServer.Error())
		http.Error(w, ErrBadServer.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"count":     len(consumers),
		"consumers": consumers,
	})
}

func (cm *ConsumerManager) RegisterConsumerHandler(w http.ResponseWriter, req *http.Request) {

	type RegisterConsumerRequest struct {
		Username string `json:"username" validate:"required"`
	}

	var body = RegisterConsumerRequest{}

	err := json.NewDecoder(req.Body).Decode(&body)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"consumer": created,
	})
}
