package admin

import (
	"encoding/json"
	"net/http"
	"time"
)

type GatewayManager struct {
}

type GatewayConfig struct {
	GatewayName string
	HostName    string
	DateCreated time.Time
	TagLine     string
}

func NewGatewayManager() *GatewayManager {
	return &GatewayManager{}
}

func LoadInitialGatewayInfo() *GatewayConfig {
	return &GatewayConfig{}
}

func (gm *GatewayManager) GetGatewayInformationHandler(w http.ResponseWriter, req *http.Request) {

	/*
		{
			"name":
			"hostName"
			"dateCreated":
			"tagline":
		}
	*/
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("all gateway information")
}
