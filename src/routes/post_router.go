package routes

import (
	"http-server/src/config"
	"http-server/src/controllers"
	"http-server/src/models"
	"log"

	"github.com/gorilla/mux"
)

func PostRouter(router *mux.Router) {
	db, err := config.Setup()

	router.HandleFunc("/posts", controllers.GetAllPosts(db, []models.Post{})).Methods("GET")
	router.HandleFunc("/post/{postID}", controllers.GetPostById(db)).Methods("GET")
	router.HandleFunc("/post", controllers.CreatePost(db, models.Post{})).Methods("POST")
	router.HandleFunc("/post/{postID}", controllers.UpdatePostById(db)).Methods("PUT")
	router.HandleFunc("/post/{postID}", controllers.DeletePostById(db)).Methods("DELETE")

	if err != nil {
		log.Panic(err)
		return
	}
}
