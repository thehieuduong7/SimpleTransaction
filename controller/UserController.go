package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"internBE.com/dto"
	models "internBE.com/entity"
	"internBE.com/helper"
	"internBE.com/service"
)

type UserController interface {
	CreateUsers(context *gin.Context)
	UpdateUsers(context *gin.Context)
	DeleteUserById(context *gin.Context)
	GetAllUsers(context *gin.Context)
	GetUserByID(context *gin.Context)
}

type userController struct {
	userService service.UserService
}

func (controller *userController) CreateUsers(context *gin.Context) {
	//TODO implement me
	var userCreateDTO dto.UserCreateDTO
	err := context.ShouldBind(&userCreateDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.JSONP(400, res)
		return
	}
	user := controller.userService.CreateUsersService(userCreateDTO)
	res := helper.BuildResponse(true, "Successful!", user)
	context.JSONP(200, res)
}

func (controller *userController) UpdateUsers(context *gin.Context) {
	//TODO implement me
	//panic("implement me")
	var userUpdateDTO dto.UserUpdateDTO
	err := context.ShouldBind(&userUpdateDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to process", err.Error(), helper.EmptyObj{})
		context.JSONP(400, res)
		return
	}
	//userUpdateDTO.ID = userId
	user := controller.userService.UpdateUsersService(userUpdateDTO)
	res := helper.BuildResponse(true, "Successful!", user)
	context.JSONP(200, res)
}

func (controller *userController) DeleteUserById(context *gin.Context) {
	//TODO implement me
	//var user models.User
	userId, err := strconv.ParseUint(context.Param("user_id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	//user.UserId = id
	controller.userService.DeleteUserService(int(userId))
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}

func (controller *userController) GetAllUsers(context *gin.Context) {
	//TODO implement me
	var users []models.User = controller.userService.GetAllUsersService()
	res := helper.BuildResponse(true, "Successful", users)
	context.JSONP(202, res)
	//panic("implement me")
}

func (controller *userController) GetUserByID(context *gin.Context) {
	//TODO implement me
	userId := context.Param("user_id")
	newUserId, err := strconv.Atoi(userId)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	user := controller.userService.GetUserByIDService(int(newUserId))

	res := helper.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)

}

func NewUserController(userServ service.UserService) UserController {

	return &userController{userService: userServ}
}
