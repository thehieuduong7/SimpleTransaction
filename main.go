package main

import (
	"github.com/gin-gonic/gin"

	"internBE.com/routes"
	"internBE.com/storage"
)

func main() {

	router := gin.New()
	storage.Connect()
	routes.AccountRoute(router)
	routes.UserRoute(router)
	router.Run(":8000")

}
