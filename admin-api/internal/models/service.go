package models

import "time"

type Service struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	TargetUrl string    `json:"targetUrl"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
