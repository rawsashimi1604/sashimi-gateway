package models

import "time"

// "required" field used in json schema validation during POST requests
type Service struct {
	Id          int       `json:"id"`
	Name        string    `json:"name" validate:"required"`
	TargetUrl   string    `json:"targetUrl" validate:"required"`
	Path        string    `json:"path" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Routes      []Route   `json:"routes"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
