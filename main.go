package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"

	"github.com/joho/godotenv"
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

	router := gin.New()
	storage.Connect(config)

	// var DB *gorm.DB
	// var test repository.UserRepository
	// test = repository.NewUserRepository(DB)
	// user := models.User{Name: "hello", PhoneNumber: "1231441", Email: "hello@example.com"}

	// test.CreateUsers(user)

	// call router
	// routes.UserRoute(router)
	router.Run(":8000")

}
