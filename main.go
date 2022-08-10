package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	models "internBE.com/model"
	"internBE.com/respository"
	"internBE.com/storage"
)

func main() {
	err := godotenv.Load(".env")

	config := &storage.Config{
		Host:     os.Getenv("DBHost"),
		Port:     os.Getenv("DBPort"),
		Password: os.Getenv("DBPassword"),
		User:     os.Getenv("DBUser"),
		DBName:   os.Getenv("DBName")}

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	type CreateBook struct {
		Name string `json:"name" binding:"required"`
	}

	router := gin.New()
	storage.Connect(config)

	user := &models.User{Name: "hello", PhoneNumber: "1231441", Email: "hello@example.com"}
	respository.UserRepository.CreateBook(user)

	// call router
	// routes.UserRoute(router)
	router.Run(":8000")

}
