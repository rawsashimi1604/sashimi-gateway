package controllers

import (
	"encoding/json"
	"net/http"
)

type MockController struct{}

func (mc *MockController) HandleIndex(w http.ResponseWriter, req *http.Request) {

	// Some implementation
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"mock": "some mock response!",
	})
}
