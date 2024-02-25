package main

import (
	"fmt"
	"http-server/src/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

	fmt.Printf("Listening at http://localhost:8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//   w.Header().Set("Content-Type", "application/json")

//   json.NewEncoder(w).Encode(map[string]string{"data": "hello world"})
// }).Methods("GET")
