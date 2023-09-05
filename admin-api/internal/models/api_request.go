package models

import (
	"time"

	"github.com/google/uuid"
)

// TODO: refactor ApiRequest to use UUID.
type ApiRequest struct {
	Id        uuid.UUID `json:"id"`
	ServiceId int       `json:"serviceId"`
	RouteId   int       `json:"routeId"`
	Path      string    `json:"path"`
	Method    string    `json:"method"`
	Time      time.Time `json:"time"`
	Code      int       `json:"code"`
}

type AggregatedApiRequest struct {
	TimeBucket time.Time `json:"timeBucket"`
	Count      int       `json:"count"`
	Count2xx   int       `json:"count_2xx"`
	Count4xx   int       `json:"count_4xx"`
	Count5xx   int       `json:"count_5xx"`
}
