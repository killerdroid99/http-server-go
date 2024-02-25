package routes

import (
	"http-server/src/controllers"

	"github.com/gorilla/mux"
)

func MeRoute(router *mux.Router) {
	router.HandleFunc("/me", controllers.Me()).Methods("GET")
}
