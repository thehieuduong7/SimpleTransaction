package controller

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUsers(context *gin.Context)
	UpdateUsers(context *gin.Context)
	DeleteUserById(context *gin.Context)
	GetAllUsers(context *gin.Context)
	GetUserByID(context *gin.Context)
}
