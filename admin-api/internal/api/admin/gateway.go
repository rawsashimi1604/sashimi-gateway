package admin

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/config"
)

type GatewayManager struct {
	GatewayConfig *GatewayConfig
}

type GatewayConfig struct {
	GatewayName string    `json:"gatewayName"`
	HostName    string    `json:"hostName"`
	DateCreated time.Time `json:"dateCreated"`
	TagLine     string    `json:"tagLine"`
	Port        string    `json:"port"`
}

func NewGatewayManager(gc *GatewayConfig) *GatewayManager {
	return &GatewayManager{
		GatewayConfig: gc,
	}
}

func LoadInitialGatewayInfo(env config.EnvVars) *GatewayConfig {
	return &GatewayConfig{
		GatewayName: env.SASHIMI_GATEWAY_NAME,
		HostName:    env.SASHIMI_HOSTNAME,
		DateCreated: time.Now(),
		TagLine:     env.SASHIMI_TAGLINE,
		Port:        env.SASHIMI_LOCAL_PORT,
	}
}

func (gm *GatewayManager) GetGatewayInformationHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"gateway": gm.GatewayConfig,
	})
}
