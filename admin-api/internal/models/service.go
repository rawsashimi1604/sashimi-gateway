package models

import "time"

type Service struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	TargetUrl   string    `json:"targetUrl"`
	Path        string    `json:"path"`
	Description string    `json:"description"`
	Routes      []Route   `json:"routes"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
