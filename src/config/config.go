package config

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	Host     = "localhost"
	User     = "postgres"
	Password = "postgres"
	Name     = "robo"
	Port     = "5432"
)

func Setup() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		Host,
		Port,
		User,
		Name,
		Password,
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func RedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}
