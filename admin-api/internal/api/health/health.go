package health

import (
	"net/http"
	"sync"

	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/service"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/utils"
	"github.com/rs/zerolog/log"
)

// TODO: implement health checks
/*
	Should change database schema for service, is_health_check_enabled, as well as health (not_enabled, startup, healthy, unhealthy)
	Cron job should run based on specified health check interval sest by gateway config
	query the /healthz route async using goroutines, see if we get a 200 OK response
	if 200 OK, then, update the health check tables in the services.
*/

type HealthChecker struct {
	serviceHealthMap map[int]string
	serviceGateway   service.ServiceGateway
	mutex            sync.Mutex
}

func NewHealthChecker(sg service.ServiceGateway) *HealthChecker {
	services, err := sg.GetAllServices()
	if err != nil {
		log.Panic().Msg("unable to create health checker. check database connection")
	}

	healthChecker := &HealthChecker{
		serviceHealthMap: make(map[int]string),
		serviceGateway:   sg,
		mutex:            sync.Mutex{},
	}

	// Initialize health checker.
	// When service is added, health check needs to be updated as well..
	for _, service := range services {
		healthChecker.serviceHealthMap[service.Id] = service.Health
	}

	return healthChecker
}

func (hc *HealthChecker) AddService(id int, health string) {
	hc.mutex.Lock()
	defer hc.mutex.Unlock()
	hc.serviceHealthMap[id] = health
}

func (hc *HealthChecker) DeleteService(id int) {
	hc.mutex.Lock()
	defer hc.mutex.Unlock()
	delete(hc.serviceHealthMap, id)

}

func (hc *HealthChecker) UpdateServiceHealth(id int, health string) {
	hc.mutex.Lock()
	defer hc.mutex.Unlock()
	hc.serviceHealthMap[id] = health
}

func (hc *HealthChecker) PingAllServices() {
	services, err := hc.serviceGateway.GetAllServices()
	if err != nil {
		log.Info().Msg("unable to check service health. cant get service information from database.")
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(services))

	for _, service := range services {
		s := service // Have to redeclare variable or goroutine wont work async
		go func(s *models.Service) {
			defer wg.Done()
			if !s.HealthCheckEnabled {
				return
			}
			res, err := http.Get(s.TargetUrl + "/healthz")
			if err != nil {
				log.Info().Msg("something went wrong, unable to ping /healthz route for service: " + s.Name)
				hc.UpdateServiceHealth(s.Id, "unhealthy")
				return
			}
			if res.StatusCode == http.StatusOK {
				hc.UpdateServiceHealth(s.Id, "healthy")
			} else {
				hc.UpdateServiceHealth(s.Id, "unhealthy")
			}
		}(&s)
	}

	wg.Wait()
	log.Info().Msg("health: " + utils.JSONStringify(hc.serviceHealthMap))
}
