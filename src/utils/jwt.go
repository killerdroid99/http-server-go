package utils

import (
	"http-server/src/structs"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GenerateJWT(userId uuid.UUID, userName string) (string, error) {
	signingKey := []byte("AllYourBase")

	claims := structs.UserClaims{
		UserID:   userId,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(signingKey)

	return ss, err
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &structs.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	return token, err
}
