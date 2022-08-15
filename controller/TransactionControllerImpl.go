package controller

import (
	"github.com/gin-gonic/gin"
	"internBE.com/dto"
	response "internBE.com/respone"
	"internBE.com/service"
)

type transactionControllerImpl struct {
	transactionService service.TransactionService
}

func NewTransactionControllerImpl(transactionService *service.TransactionService) TransactionController {
	return &transactionControllerImpl{transactionService: *transactionService}
}

func (t *transactionControllerImpl) Create(ctx *gin.Context) {
	var reqTransaction dto.CreateTransactionRequest
	err := ctx.ShouldBindJSON(&reqTransaction)
	if err != nil {
		response.Fail(ctx, 500, err.Error())
		return
	}
	resTransaction, err := t.transactionService.Create(reqTransaction)
	if err != nil {
		response.Fail(ctx, 500, err.Error())
		return
	}
	response.Success(ctx, "data success", resTransaction)
}
func (t *transactionControllerImpl) GetAllTransRelatedNumberAcc(ctx *gin.Context) {
	var reqTransaction dto.GetMyTransactionRequest
	err := ctx.ShouldBindJSON(&reqTransaction)
	if err != nil {
		response.Fail(ctx, 500, err.Error())
		return
	}
	resTransactions, err := t.transactionService.GetAllTransRelatedNumberAcc(reqTransaction)
	if err != nil {
		response.Fail(ctx, 500, err.Error())
		return
	}
	response.Success(ctx, "GetAllTransRelatedNumberAcc", resTransactions)
}
func (t *transactionControllerImpl) GetTransSendedByNumberAcc(ctx *gin.Context) {
	var reqTransaction dto.GetMyTransactionRequest
	err := ctx.ShouldBindJSON(&reqTransaction)
	if err != nil {
		response.Fail(ctx, 500, err.Error())
		return
	}
	resTransactions, err := t.transactionService.GetTransSendedByNumberAcc(reqTransaction)
	if err != nil {
		response.Fail(ctx, 500, err.Error())
		return
	}
	response.Success(ctx, "GetTransSendedByNumberAcc", resTransactions)
}
func (t *transactionControllerImpl) GetTransRevievedByNumberAcc(ctx *gin.Context) {
	var reqTransaction dto.GetMyTransactionRequest
	err := ctx.ShouldBindJSON(&reqTransaction)
	if err != nil {
		response.Fail(ctx, 500, err.Error())
		return
	}
	resTransactions, err := t.transactionService.GetTransRevievedByNumberAcc(reqTransaction)
	if err != nil {
		response.Fail(ctx, 500, err.Error())
		return
	}
	response.Success(ctx, "GetTransRevievedByNumberAcc", resTransactions)
}
func (t *transactionControllerImpl) GetTransFromTo(ctx *gin.Context) {
	var reqTransaction dto.GetTransactionFromToRequest
	err := ctx.ShouldBindJSON(&reqTransaction)
	if err != nil {
		response.Fail(ctx, 500, err.Error())
		return
	}
	resTransactions, err := t.transactionService.GetTransFromTo(reqTransaction)
	if err != nil {
		response.Fail(ctx, 500, err.Error())
		return
	}
	response.Success(ctx, "GetTransFromTo", resTransactions)
}
func (t *transactionControllerImpl) GetTransDateToDate(ctx *gin.Context) {
	var reqTransaction dto.GetTransactionFromToTimeRequest
	err := ctx.ShouldBindJSON(&reqTransaction)
	if err != nil {
		response.Fail(ctx, 500, err.Error())
		return
	}
	resTransactions, err := t.transactionService.GetTransDateToDate(reqTransaction)
	if err != nil {
		response.Fail(ctx, 500, err.Error())
		return
	}
	response.Success(ctx, "GetTransDateToDate", resTransactions)
}