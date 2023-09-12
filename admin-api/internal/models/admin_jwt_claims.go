package models

import "github.com/golang-jwt/jwt/v5"

type AdminJWTClaims struct {
	Gateway string `json:"gateway"`
	jwt.RegisteredClaims
}
