package jobs

import (
	"time"

	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api/analytics"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/gateway/request"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

// TODO: create gateway for adding requests.
type RequestCronJob struct {
	AnalyticsTracker *analytics.AnalyticsTracker
	RequestGateway   *request.RequestGateway
	Cron             *cron.Cron
	Interval         time.Duration
}

func NewRequestCronJob(at *analytics.AnalyticsTracker, rg *request.RequestGateway, interval time.Duration) *RequestCronJob {
	return &RequestCronJob{
		AnalyticsTracker: at,
		RequestGateway:   rg,
		Cron:             cron.New(),
		Interval:         interval,
	}
}

func (rcj *RequestCronJob) Start() {
	rcj.Cron.AddFunc("@every "+rcj.Interval.String(), func() {
		rcj.Run()
	})
	rcj.Cron.Start()
}

func (rcj *RequestCronJob) Run() {
	log.Info().Msg("Ran the cron job!")
}
