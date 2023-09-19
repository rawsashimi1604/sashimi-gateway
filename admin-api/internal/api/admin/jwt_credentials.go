package admin

import (
	"encoding/json"
	"net/http"

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
