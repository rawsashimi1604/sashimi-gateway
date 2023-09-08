package jobs

import (
	"time"

	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api/analytics"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/websocket"
	"github.com/robfig/cron/v3"
)

// TODO: create gateway for adding requests.
type RequestCronJob struct {
	AnalyticsTracker *analytics.AnalyticsTracker
	Cron             *cron.Cron
	Interval         time.Duration
	WebsocketServer  *websocket.WebSocketServer
}

func NewRequestCronJob(at *analytics.AnalyticsTracker, interval time.Duration, ws *websocket.WebSocketServer) *RequestCronJob {
	return &RequestCronJob{
		AnalyticsTracker: at,
		Cron:             cron.New(),
		Interval:         interval,
		WebsocketServer:  ws,
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
	rcj.WebsocketServer.BroadcastMessage([]byte("hello world"))
}
