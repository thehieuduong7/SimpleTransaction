package main

import (
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

	// router := gin.New()
	storage.Connect(config)
	// call router
	// UserRoute(router)
	// router.Run(":8000")

}
