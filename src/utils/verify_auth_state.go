package utils

import (
	"net/http"
)

func VerifyAuthState(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("JWT")

	if err != nil {
		return false
	}

	if _, err = VerifyJWT(cookie.Value); err != nil {
		return false
	}

	return true
}
