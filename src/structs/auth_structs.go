package structs

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type LoginUser struct {
	Email    string
	Password string
}

type RegisterUser struct {
	Name          string
	Email         string
	Password      string
	LoginDirectly bool
}

type UserClaims struct {
	jwt.RegisteredClaims
	UserID   uuid.UUID `json:"id"`
	UserName string    `json:"name"`
}
