package controller

import (
	response "internBE.com/respone"
	"strconv"

	"github.com/gin-gonic/gin"
	"internBE.com/dto"
	"internBE.com/service"
)

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
		response.Fail(context, 500, "invalid format string")
		return
	}
	err1 := controller.AccountService.CreateAccount(&accountDto)

	if err1 != nil {
		response.Fail(context, 500, "the query fail")
	}
	response.Success(context, "successfully", accountDto)
}

func (controller *accountController) UpdateAccount(context *gin.Context) {
	var accountDto = dto.AccountDtoToInsert{}
	err := context.ShouldBind(&accountDto)
	if err != nil {
		response.Fail(context, 500, "invalid format string")
	}
	err1 := controller.AccountService.UpdateAccount(&accountDto)
	if err1 != nil {
		response.Fail(context, 500, "the query fail")
	}
	response.Success(context, "successfully", accountDto)
}

func (controller *accountController) DeleteAccount(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)
	if err != nil {
		response.Fail(context, 500, "invalid format string")
	}
	err1 := controller.AccountService.DeleteAccount(int(id))
	if err1 != nil {
		response.Fail(context, 500, "the query fail")
	}
	response.Success(context, "successfully", nil)
}

func (controller *accountController) GetAccountByUserId(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 0, 0)
	if err != nil {
		response.Fail(context, 500, "invalid format string")
	}
	var account, err1 = controller.AccountService.GetAccountByUserId(int(id))
	if err1 != nil {
		response.Fail(context, 500, "not found")
	} else {
		response.Success(context, "successfully", account)
	}
}

func (controller *accountController) GetAccountById(context *gin.Context) {
	id, err1 := strconv.ParseInt(context.Param("id"), 0, 0)
	if err1 != nil {
		response.Fail(context, 500, "invalid format string")
	}
	var account, err = controller.AccountService.GetAccountById(int(id))
	if err != nil {
		response.Fail(context, 500, "not found")
	} else {
		response.Success(context, "successfully", account)
	}

}
