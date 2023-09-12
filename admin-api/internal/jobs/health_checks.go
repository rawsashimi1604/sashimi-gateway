package jobs

import (
	"time"

	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api/health"
	"github.com/robfig/cron/v3"
)

type HealthCheckCronJob struct {
	HealthChecker *health.HealthChecker
	Cron          *cron.Cron
	Interval      time.Duration
}

func NewHealthCheckCronJob(
	hc *health.HealthChecker,
	interval time.Duration,
) *HealthCheckCronJob {
	return &HealthCheckCronJob{
		HealthChecker: hc,
		Cron:          cron.New(),
		Interval:      interval,
	}
}

func (hcj *HealthCheckCronJob) Start() {
	hcj.Cron.AddFunc("@every "+hcj.Interval.String(), func() {
		hcj.run()
	})
	hcj.Cron.Start()
}

func (hcj *HealthCheckCronJob) run() {
	hcj.HealthChecker.PingAllServices()
}
