package routes

import (
	"http-server/src/config"
	"http-server/src/controllers"
	"log"

	"github.com/gorilla/mux"
)

func LoginRoute(router *mux.Router) {
	db, err := config.Setup()

	router.HandleFunc("/login", controllers.LoginUser(db)).Methods("POST")

	if err != nil {
		log.Panic(err)
		return
	}
}

func LogoutRoute(router *mux.Router) {
	router.HandleFunc("/logout", controllers.LogoutUser()).Methods("POST")
}

func MeRoute(router *mux.Router) {
	router.HandleFunc("/me", controllers.Me()).Methods("GET")
}

func RegisterRoute(router *mux.Router) {
	db, err := config.Setup()

	router.HandleFunc("/register", controllers.RegisterUser(db)).Methods("POST")

	if err != nil {
		log.Panic(err)
		return
	}
}
