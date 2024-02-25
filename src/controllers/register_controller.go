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

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB, user models.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		registerUserInput := structs.RegisterUser{}

		// check if session already exists
		if utils.VerifyAuthState(w, r) {
			utils.SetResponse(w, http.StatusNotAcceptable, "error", "Session already exists")
			return
		}

		// validate request body
		if err := json.NewDecoder(r.Body).Decode(&registerUserInput); err != nil {
			utils.SetResponse(w, http.StatusBadRequest, "error", err.Error())
			return
		}

		// trimming input data for blank spaces
		name := strings.Trim(registerUserInput.Name, " ")
		email := strings.Trim(registerUserInput.Email, " ")
		password := strings.Trim(registerUserInput.Password, " ")

		// basic input validation
		if len(name) < 3 {
			utils.SetResponse(w, http.StatusBadRequest, "error", map[string]string{"value": "Name must be at least 3 characters long", "field": "name"})
			return
		}
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
		if existingUser.Error == nil {
			utils.SetResponse(w, http.StatusConflict, "error", fmt.Sprintf("User with email: %s already exists", email))
			return
		}

		hashedPwd, err := utils.HashPassword(password)

		if err != nil {
			utils.SetResponse(w, http.StatusInternalServerError, "error", err.Error())
			return
		}
		newUser := models.User{
			ID:       uuid.New(),
			Name:     name,
			Email:    email,
			Password: hashedPwd,
		}

		db.Create(&newUser)

		if registerUserInput.LoginDirectly {
			ss, err := utils.GenerateJWT(newUser.ID, newUser.Name)

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

				utils.SetResponse(w, http.StatusCreated, "success", map[string]any{
					"msg": "User successfully registered and authenticated",
					"user": map[string]any{
						"id":        newUser.ID,
						"name":      newUser.Name,
						"email":     newUser.Email,
						"createdAt": newUser.CreatedAt,
					},
				})
				return
			}
		}

		utils.SetResponse(w, http.StatusCreated, "success", map[string]any{
			"msg": "User successfully registered",
			"user": map[string]any{
				"id":        newUser.ID,
				"name":      newUser.Name,
				"email":     newUser.Email,
				"createdAt": newUser.CreatedAt,
			},
		})
	}
}
