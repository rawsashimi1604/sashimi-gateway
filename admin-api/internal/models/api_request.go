package models

import (
	"time"
)

// TODO: refactor ApiRequest to use UUID.
type ApiRequest struct {
	Id        int
	ServiceId int
	RouteId   int
	Path      string
	Method    string
	Time      time.Time
	Code      int
}
