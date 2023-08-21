package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

var (
	ErrReadHttpBodyFail = errors.New("unable to read http body")
)

type JSONErrorMessage struct {
	Message string `json:"message"`
}

func JSONError(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(JSONErrorMessage{
		Message: err,
	})
}

func JSONStringify(content any) string {
	jsonified, _ := json.Marshal(content)
	return string(jsonified)
}

func ReadHttpBody(body io.ReadCloser) ([]byte, error) {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		log.Info().Msg("Failed to read http body.")
		return nil, ErrReadHttpBodyFail
	}
	return bodyBytes, nil
}
