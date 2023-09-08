package jobs

import (
	"time"

	socketio "github.com/googollee/go-socket.io"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api/analytics"
	"github.com/robfig/cron/v3"
)

// TODO: create gateway for adding requests.
type RequestCronJob struct {
	AnalyticsTracker *analytics.AnalyticsTracker
	Cron             *cron.Cron
	Interval         time.Duration
	Server           *socketio.Server
}

func NewRequestCronJob(at *analytics.AnalyticsTracker, interval time.Duration, server *socketio.Server) *RequestCronJob {
	return &RequestCronJob{
		AnalyticsTracker: at,
		Cron:             cron.New(),
		Interval:         interval,
		Server:           server,
	}
}

func (rcj *RequestCronJob) Start() {
	rcj.Cron.AddFunc("@every "+rcj.Interval.String(), func() {
		rcj.run()
	})
	rcj.Cron.Start()
}

func (rcj *RequestCronJob) run() {
	requests := rcj.AnalyticsTracker.StoreRequests()
	rcj.Server.BroadcastToRoom("", "bcast", "event:apiRequests", map[string]interface{}{
		"requests": requests,
	})
}
