package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	var passwordBytes = []byte(password)

	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err
}

func VerifyPassword(hash string, password string) bool {
	var passwordBytes = []byte(password)
	var hashBytes = []byte(hash)

	if err := bcrypt.CompareHashAndPassword(hashBytes, passwordBytes); err != nil {
		return false
	}

	return true
}
