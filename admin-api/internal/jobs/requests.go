package jobs

import (
	"time"

	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api/analytics"
	"github.com/robfig/cron/v3"
)

// TODO: create gateway for adding requests.
type RequestCronJob struct {
	AnalyticsTracker *analytics.AnalyticsTracker
	Cron             *cron.Cron
	Interval         time.Duration
}

func NewRequestCronJob(at *analytics.AnalyticsTracker, interval time.Duration) *RequestCronJob {
	return &RequestCronJob{
		AnalyticsTracker: at,
		Cron:             cron.New(),
		Interval:         interval,
	}
}

func (rcj *RequestCronJob) Start() {
	rcj.Cron.AddFunc("@every "+rcj.Interval.String(), func() {
		rcj.run()
	})
	rcj.Cron.Start()
}

func (rcj *RequestCronJob) run() {
	rcj.AnalyticsTracker.StoreRequests()
}
