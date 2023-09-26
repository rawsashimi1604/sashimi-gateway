package admin

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/consumer"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/jwt_credentials"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/utils"
	"github.com/rs/zerolog/log"
)

var (
	ErrInvalidCredentialBody = errors.New("invalid jwt credentials body")
)

type JwtCredentialsManager struct {
	JwtCredentialsGateway jwt_credentials.JWTCredentialsGateway
	ConsumerGateway       consumer.ConsumerGateway
}

func NewJwtCredentialsManager(jcg jwt_credentials.JWTCredentialsGateway, cg consumer.ConsumerGateway) *JwtCredentialsManager {
	return &JwtCredentialsManager{
		JwtCredentialsGateway: jcg,
		ConsumerGateway:       cg,
	}
}

func (jcm *JwtCredentialsManager) CreateNewCredentialHandler(w http.ResponseWriter, req *http.Request) {

	type CreateNewCredentialRequest struct {
		ConsumerId string `json:"consumerId" validate:"required"`
		Name       string `json:"name" validate:"required"`
	}

	var body = CreateNewCredentialRequest{}

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Info().Msg(ErrInvalidCredentialBody.Error())
		http.Error(w, ErrInvalidCredentialBody.Error(), http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(body.ConsumerId)
	if err != nil {
		log.Info().Msg("invalid id passed")
		http.Error(w, "invalid id passed", http.StatusBadRequest)
		return
	}

	cmer, err := jcm.ConsumerGateway.GetConsumerById(id)
	if err != nil {
		if err == consumer.ErrConsumerNotFound {
			log.Info().Msg("consumer does not exist")
			http.Error(w, "consumer does not exist", http.StatusNotFound)
			return
		}
		log.Info().Msg(err.Error())
		log.Info().Msg("something went wrong when finding consumer")
		http.Error(w, "something went wrong when finding consumer", http.StatusInternalServerError)
		return
	}

	jwtKey, _ := utils.GenerateRandomKey()
	jwtSecret, _ := utils.GenerateRandomKey()

	credential := models.JWTCredentials{
		Id:        uuid.New(),
		Key:       jwtKey,
		Secret:    jwtSecret,
		Name:      body.Name,
		Consumer:  cmer,
		CreatedAt: time.Now(),
	}

	created, err := jcm.JwtCredentialsGateway.AddCredential(credential)
	if err != nil {
		log.Info().Msg("something went wrong when adding jwt_credentials")
		http.Error(w, "something went wrong when adding jwt_credentials", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"jwt_credentials": created,
	})

}

func (jcm *JwtCredentialsManager) ListCredentialsHandler(w http.ResponseWriter, req *http.Request) {
	credentials, err := jcm.JwtCredentialsGateway.ListCredentials()
	if err != nil {
		log.Info().Msg(err.Error())
		http.Error(w, "error retrieving credentials", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"count":           len(credentials),
		"jwt_credentials": credentials,
	})
}

func (jcm *JwtCredentialsManager) GetAllCredentialsByConsumerId(w http.ResponseWriter, req *http.Request) {
	consumerId := mux.Vars(req)["id"]
	consumerIdConverted, err := uuid.Parse(consumerId)
	if err != nil {
		log.Info().Msg(err.Error())
		http.Error(w, "invalid id passed in, must be uuid", http.StatusBadRequest)
		return
	}

	credentials, err := jcm.JwtCredentialsGateway.GetAllCredentialsByConsumer(consumerIdConverted)
	if err != nil {
		log.Info().Msg(err.Error())
		http.Error(w, "error retrieving credentials", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"count":           len(credentials),
		"jwt_credentials": credentials,
	})
}
