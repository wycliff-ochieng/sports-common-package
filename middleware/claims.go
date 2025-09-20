package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	ID     int       `json:"id"`
	UserID uuid.UUID `json:"userid"`
	Roles  []string  `json:"roles"`
	Email  string    `json:"email"`
	jwt.RegisteredClaims
}
