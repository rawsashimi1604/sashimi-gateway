package models

import "time"

type ApiRequest struct {
	Id        int
	ServiceId int
	RouteId   int
	Path      string
	Method    string
	Time      time.Time
}
