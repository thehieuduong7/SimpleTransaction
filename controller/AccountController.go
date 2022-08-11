package controller

import (
	"github.com/gin-gonic/gin"
	"internBE.com/dto"
	"internBE.com/helper"
	"internBE.com/service"
	"net/http"
	"strconv"
)

type AccountController interface {
	CreateAccount(context *gin.Context)
	UpdateAccount(context *gin.Context)
	DeleteAccount(context *gin.Context)
	GetAccountByUserId(context *gin.Context)
	GetAccountById(context *gin.Context)
}
type accountController struct {
	AccountService service.AccountService
}

func NewAccountController(accountService service.AccountService) AccountController {
	return &accountController{AccountService: accountService}
}

func (controller *accountController) CreateAccount(context *gin.Context) {
	var accountDto dto.AccountDtoToInsert
	err := context.ShouldBindJSON(&accountDto)
	if err != nil {
		res := helper.Response{Message: "fail to get body request", Status: false, Errors: err, Data: helper.EmptyOb{}}
		context.JSON(http.StatusBadRequest, res)
		return
	}
	controller.AccountService.CreateAccount(&accountDto)
	res := helper.Response{Message: "oki", Status: true, Errors: nil, Data: accountDto}
	context.JSON(http.StatusCreated, res)
}

func (controller *accountController) UpdateAccount(context *gin.Context) {
	var accountDto = dto.AccountDtoToInsert{}
	err := context.ShouldBind(&accountDto)
	if err != nil {
		res := helper.Response{Message: "fail to get body request", Status: false, Errors: err, Data: helper.EmptyOb{}}
		context.JSON(http.StatusBadRequest, res)
	}
	controller.AccountService.UpdateAccount(&accountDto)
	res := helper.Response{Message: "oki", Status: true, Errors: nil, Data: accountDto}
	context.JSON(http.StatusCreated, res)
}

func (controller *accountController) DeleteAccount(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.Response{Message: "fail to get parameter request", Status: false, Errors: err, Data: helper.EmptyOb{}}
		context.JSON(http.StatusBadRequest, res)
	}
	controller.AccountService.DeleteAccount(int(id))
	context.JSON(http.StatusOK, "oki")
}

func (controller *accountController) GetAccountByUserId(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.Response{Message: "fail to get parameter request", Status: false, Errors: err, Data: helper.EmptyOb{}}
		context.JSON(http.StatusBadRequest, res)
	}
	var account, err1 = controller.AccountService.GetAccountByUserId(int(id))
	if err1 != nil {
		res := helper.Response{Message: "not found account", Status: false, Errors: err, Data: helper.EmptyOb{}}
		context.JSON(http.StatusBadRequest, res)

	} else {
		res := helper.Response{Message: "oki", Status: true, Errors: nil, Data: account}
		context.JSON(http.StatusOK, res)
	}

}

func (controller *accountController) GetAccountById(context *gin.Context) {
	id, err1 := strconv.ParseInt(context.Param("id"), 0, 0)
	if err1 != nil {
		res := helper.Response{Message: "fail to get parameter request", Status: false, Errors: err1, Data: helper.EmptyOb{}}
		context.JSON(http.StatusBadRequest, res)
	}
	var account, err = controller.AccountService.GetAccountById(int(id))
	if err != nil {
		res := helper.Response{Message: "not found account", Status: false, Errors: err, Data: helper.EmptyOb{}}
		context.JSON(http.StatusBadRequest, res)
	} else {
		res := helper.Response{Message: "oki", Status: true, Errors: nil, Data: account}
		context.JSON(http.StatusOK, res)
	}

}
