package utils

import (
	"encoding/json"
	"net/http"
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
