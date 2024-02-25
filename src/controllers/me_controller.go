package controllers

import (
	"http-server/src/structs"
	"http-server/src/utils"
	"net/http"
)

func Me() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		cookie, err := r.Cookie("JWT")

		if err != nil {
			utils.SetResponse(w, http.StatusBadRequest, "error", "Invalid or Expired token")
			return
		}

		token, err := utils.VerifyJWT(cookie.Value)

		if err != nil {
			utils.SetResponse(w, http.StatusInternalServerError, "error", err.Error())
			return
		}

		userData := map[string]any{
			"userId":   token.Claims.(*structs.UserClaims).UserID,
			"userName": token.Claims.(*structs.UserClaims).UserName,
		}

		utils.SetResponse(w, http.StatusOK, "success", userData)
	}
}
