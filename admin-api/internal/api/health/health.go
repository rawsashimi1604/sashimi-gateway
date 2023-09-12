package health

import "github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"

// TODO: implement health checks
/*
	Should change database schema for service, is_health_check_enabled, as well as health (not_enabled, startup, healthy, unhealthy)
	Cron job should run based on specified health check interval sest by gateway config
	query the /healthz route async using goroutines, see if we get a 200 OK response
	if 200 OK, then, update the health check tables in the services.
*/

type HealthChecker struct {
}

func NewHealthChecker() *HealthChecker {
	return &HealthChecker{}
}

func (hc *HealthChecker) CheckServiceHealth(service *models.Service) {}
