package main

import (
	"github.com/gin-gonic/gin"
	"internBE.com/routes"
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
	routes.AccountRoute(router)

	router.Run(":8000")

}
