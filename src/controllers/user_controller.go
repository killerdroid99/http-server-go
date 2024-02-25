package controllers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"http-server/src/models"
// 	"http-server/src/responses"
// 	"http-server/src/utils"
// 	"net/http"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// var validate = validator.New()

// func CreateUser(db *gorm.DB, payment models.User) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		user := models.User{}

// 		// validate request body
// 		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			response := responses.UserResponses{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
// 			json.NewEncoder(w).Encode(response)
// 			return
// 		}

// 		if validationErr := validate.Struct(&user); validationErr != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			response := responses.UserResponses{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
// 			json.NewEncoder(w).Encode(response)
// 			return
// 		}

// 		hashedPwd, err := utils.HashPassword(user.Password)

// 		if err == nil {
// 			newUser := models.User{
// 				ID:       uuid.New(),
// 				Name:     user.Name,
// 				Email:    user.Email,
// 				Password: hashedPwd,
// 			}

// 			result := db.Create(&newUser)
// 			fmt.Println(result)
// 		}
// 	}
// }
