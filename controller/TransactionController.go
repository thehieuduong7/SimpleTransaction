package controller

import "github.com/gin-gonic/gin"

type TransactionController interface {
	Create(ctx *gin.Context)
	GetAllTransRelatedNumberAcc(ctx *gin.Context)
	GetHistoryAccountNo(ctx *gin.Context)
	GetTransFromTo(ctx *gin.Context)
	GetTransDateToDate(ctx *gin.Context)
	GetStaticTransaction(ctx *gin.Context)
}
