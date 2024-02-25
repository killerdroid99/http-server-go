package controllers

import (
	"http-server/src/utils"
	"net/http"
)

func LogoutUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		cookie, err := r.Cookie("JWT")

		if err != nil {
			utils.SetResponse(w, http.StatusBadRequest, "error", "Not authenticated")
			return
		}

		cookie.MaxAge = -1

		http.SetCookie(w, cookie)
		utils.SetResponse(w, http.StatusOK, "success", "User logged out")
	}
}
