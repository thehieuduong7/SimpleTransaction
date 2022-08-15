package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"internBE.com/pkg/seeds"
	"internBE.com/routes"
	"internBE.com/storage"
)

func main() {

	router := gin.New()
	storage.GetDB()
	// GetSeedData()
	routes.AccountRoute(router)
	routes.UserRoute(router)
	routes.TransactionRoute(router)
	router.Run(":8000")
}

func GetSeedData() {
	for _, seed := range seeds.All() {
		if err := seed.Run(storage.GetDB()); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}
}
