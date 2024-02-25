package controllers

import (
	"encoding/json"
	"fmt"
	"http-server/src/models"
	"http-server/src/structs"
	"http-server/src/utils"
	"net/http"
	"strings"
	"time"

	"gorm.io/gorm"
)

func LoginUser(db *gorm.DB, user models.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		loginUserInput := structs.LoginUser{}

		// check if session already exists
		if utils.VerifyAuthState(w, r) {
			utils.SetResponse(w, http.StatusNotAcceptable, "error", "Session already exists")
			return
		}

		// validate request body
		if err := json.NewDecoder(r.Body).Decode(&loginUserInput); err != nil {
			utils.SetResponse(w, http.StatusBadRequest, "error", err.Error())
			return
		}

		// trimming input data for blank spaces
		email := strings.Trim(loginUserInput.Email, " ")
		password := strings.Trim(loginUserInput.Password, " ")

		// basic input validation
		if len(email) == 0 {
			utils.SetResponse(w, http.StatusBadRequest, "error", map[string]string{"value": "Email is required", "field": "email"})
			return
		}
		if !strings.Contains(email, "@") {
			utils.SetResponse(w, http.StatusBadRequest, "error", map[string]string{"value": "Invalid email address", "field": "email"})
			return
		}
		if len(password) < 5 {
			utils.SetResponse(w, http.StatusBadRequest, "error", map[string]string{"value": "Password must be at least 5 characters long", "field": "password"})
			return
		}

		// check if user exists
		existingUser := db.First(&user, "email = ?", email)
		if existingUser.Error != nil {
			utils.SetResponse(w, http.StatusBadRequest, "error", fmt.Sprintf("User with email: %s not found", email))
			return
		}

		// verifying password
		if validPassword := utils.VerifyPassword(user.Password, password); !validPassword {
			utils.SetResponse(w, http.StatusBadRequest, "error", "Wrong password")
			return
		}

		ss, err := utils.GenerateJWT(user.ID, user.Name)

		// set JWT to cookie
		if err == nil {
			cookie := http.Cookie{
				Name:     "JWT",
				Value:    ss,
				Path:     "/",
				MaxAge:   int(25 * time.Hour),
				HttpOnly: false,
				Secure:   true,
				SameSite: http.SameSiteLaxMode,
			}

			// setting cookie
			http.SetCookie(w, &cookie)

			utils.SetResponse(w, http.StatusOK, "success", "User authenticated")
		}
	}
}
