package models

import (
	"time"

	"github.com/google/uuid"
)

type Consumer struct {
	Id             uuid.UUID `json:"id"`
	Username       string    `json:"username"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	JwtAuthEnabled bool      `json:"isJwtAuthEnabled"`
}
