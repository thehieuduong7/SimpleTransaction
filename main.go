package main

import (
	"github.com/gin-gonic/gin"

	"internBE.com/routes"
)

func main() {

	router := gin.New()
	routes.AccountRoute(router)
	routes.UserRoute(router)
	routes.TransactionRoute(router)
	router.Run(":8000")
}
