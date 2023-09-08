package jobs

import (
	"time"

	"github.com/google/uuid"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api/analytics"
	"github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/models"
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
	requests := rcj.AnalyticsTracker.StoreRequests()
	requests = append(
		requests,
		models.ApiRequest{
			Id:        uuid.New(),
			ServiceId: 1,
			RouteId:   1,
			Path:      "/test",
			Method:    "GET",
			Time:      time.Now(),
			Code:      200,
		},
	)
	rcj.WebsocketServer.BroadcastMessage("request cron job!!!")
	if len(requests) > 0 {
		rcj.WebsocketServer.BroadcastRequests(requests)
	}
}
