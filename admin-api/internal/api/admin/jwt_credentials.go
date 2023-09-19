package admin

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/jwt_credentials"
	"github.com/rs/zerolog/log"
)

type JwtCredentialsManager struct {
	JwtCredentialsGateway jwt_credentials.JWTCredentialsGateway
}

func NewJwtCredentialsManager(jcg jwt_credentials.JWTCredentialsGateway) *JwtCredentialsManager {
	return &JwtCredentialsManager{
		JwtCredentialsGateway: jcg,
	}
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
