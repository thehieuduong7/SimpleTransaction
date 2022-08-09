package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"internBE.com/storage"

	"internBE.com/pkg/seeds"
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
	for _, seed := range seeds.All() {
		if err := seed.Run(storage.Connect(config)); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}
	// call router
	// UserRoute(router)
	router.Run(":8000")

}
