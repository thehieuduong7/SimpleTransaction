package routes

import (
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.POST("/user")
		v1.POST("/submit")
		v1.POST("/read")
	}
}
