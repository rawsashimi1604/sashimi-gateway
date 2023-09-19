package models

import (
	"github.com/google/uuid"
)

type JWTCredentials struct {
	Id       uuid.UUID `json:"id"`
	Key      string    `json:"key"`
	Secret   string    `json:"secret"`
	Name     string    `json:"name"`
	Consumer Consumer  `json:"consumer"`
}
