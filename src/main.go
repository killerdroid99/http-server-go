package main

import (
	"fmt"
	"http-server/src/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}

	r := mux.NewRouter()
	// register route
	routes.RegisterRoute(r)
	// login route
	routes.LoginRoute(r)
	// logout route
	routes.LogoutRoute(r)
	// me route
	routes.MeRoute(r)

	// create new post route
	// routes.CreatePostRoute(r)

	// post router
	routes.PostRouter(r)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:5173"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch},
		AllowCredentials: true,
		Debug:            true,
	}).Handler(r)

	fmt.Printf("Listening at http://localhost:8080\n")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
