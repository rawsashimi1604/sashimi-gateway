package admin

import (
	"encoding/json"
	"net/http"
)

type GatewayManager struct {
}

func NewGatewayManager() *GatewayManager {
	return &GatewayManager{}
}

func (gm *GatewayManager) GetGatewayInformationHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("all gateway information")
}
