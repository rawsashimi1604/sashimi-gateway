package models

import "time"

type Route struct {
	Id          int       `json:"id"`
	Path        string    `json:"path"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Method      ApiMethod `json:"method"`
}
