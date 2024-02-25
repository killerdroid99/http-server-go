package db

import (
	"fmt"
	"http-server/src/config"
	"http-server/src/models"
	"log"
)

func main() {

	// connect db
	db, err := config.Setup()
	if err != nil {
		log.Panic(err)
		return
	}
	fmt.Println("Connected")
	// db auto migrate models
	db.AutoMigrate(models.User{})
}
