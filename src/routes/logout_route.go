package routes

import (
	"http-server/src/controllers"

	"github.com/gorilla/mux"
)

func LogoutRoute(router *mux.Router) {
	router.HandleFunc("/logout", controllers.LogoutUser()).Methods("POST")
}
