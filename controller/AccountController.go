package controller

import "github.com/gin-gonic/gin"

type AccountController interface {
	CreateAccount(context *gin.Context)
	UpdateAccount(context *gin.Context)
	DeleteAccount(context *gin.Context)
	GetAccountByUserId(context *gin.Context)
	GetAccountById(context *gin.Context)
}
