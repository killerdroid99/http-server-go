package routes

import (
	"http-server/src/config"
	"http-server/src/controllers"
	"http-server/src/models"
	"log"

	"github.com/gorilla/mux"
)

func LoginRoute(router *mux.Router) {
	db, err := config.Setup()

	router.HandleFunc("/login", controllers.LoginUser(db, models.User{})).Methods("POST")

	if err != nil {
		log.Panic(err)
		return
	}
}
