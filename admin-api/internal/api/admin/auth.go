package admin

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/config"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/validator"
	"github.com/rs/zerolog/log"
)

var (
	ErrInvalidLoginBody   = errors.New("invalid login body")
	ErrInvalidCredentials = errors.New("invalid login credentials")
)

type AdminAuthManager struct {
	key []byte
}

func NewAdminAuthManager(key []byte) *AdminAuthManager {
	return &AdminAuthManager{
		key: key,
	}
}

func (am *AdminAuthManager) Login(w http.ResponseWriter, req *http.Request) {

	type LoginRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	var body = LoginRequest{}

	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		log.Info().Msg(ErrInvalidLoginBody.Error())
		http.Error(w, ErrInvalidLoginBody.Error(), http.StatusBadRequest)
		return
	}

	validator := validator.NewValidator()
	err = validator.ValidateStruct(&body)
	if err != nil {
		log.Info().Msg(ErrInvalidLoginBody.Error())
		http.Error(w, ErrInvalidLoginBody.Error(), http.StatusBadRequest)
		return
	}

	env := config.LoadEnv()
	isValidCredentials := body.Username == env.SASHIMI_ADMIN_USERNAME && body.Password == env.SASHIMI_ADMIN_PASSWORD

	if !isValidCredentials {
		log.Info().Msg(ErrInvalidCredentials.Error())
		http.Error(w, ErrInvalidCredentials.Error(), http.StatusUnauthorized)
		return
	}

	// Generate new token...
	// https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-Parse-Hmac
	claims := models.AdminJWTClaims{
		env.SASHIMI_GATEWAY_NAME,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Sashimi Gateway Admin Api",
			Subject:   "auth request",
			Audience:  []string{"manager"},
		},
	}

	signed, err := am.SignJWT(&claims)
	if err != nil {
		log.Info().Msg("something went wrong when signing jwt claims")
		http.Error(w, "something went wrong when signing jwt claims", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": signed,
	})
}

func (am *AdminAuthManager) SignJWT(claims *models.AdminJWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(am.key)
	if err != nil {
		return "", errors.New("unable to sign jwt claims")
	}
	return signed, nil
}
