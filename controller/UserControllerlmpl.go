package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"internBE.com/dto"
	models "internBE.com/entity"
	response "internBE.com/respone"
	"internBE.com/service"
)

type userController struct {
	userService service.UserService
}

func (controller *userController) CreateUsers(context *gin.Context) {
	//TODO implement me
	var userCreateDTO dto.UserCreateDTO
	err := context.ShouldBind(&userCreateDTO)
	if err != nil {
		response.Fail(context, 500, "invalid format json "+err.Error())
		return
	}

	result := controller.userService.CreateUsersService(userCreateDTO)
	response.Success(context, "successfully", result)
	// context.JSON(http.StatusCreated, response)
}

func (controller *userController) UpdateUsers(context *gin.Context) {

	var userUpdateDTO dto.UserUpdateDTO
	err := context.ShouldBind(&userUpdateDTO)
	if err != nil {
		response.Fail(context, 500, "fail to get json"+err.Error())
		return
	}
	controller.userService.UpdateUsersService(userUpdateDTO)
	response.Success(context, "successfully", nil)
}

func (controller *userController) DeleteUserById(context *gin.Context) {
	//TODO implement me
	//var user models.User
	userId, err := strconv.ParseUint(context.Param("user_id"), 0, 0)
	if err != nil {
		response.Fail(context, 500, "invalid format json "+err.Error())
		return
	}
	//user.UserId = id
	controller.userService.DeleteUserService(int(userId))
	response.Success(context, "successfully", nil)
}

func (controller *userController) GetAllUsers(context *gin.Context) {
	//TODO implement me
	var users []models.User = controller.userService.GetAllUsersService()
	response.Success(context, "successfully", users)
}

func (controller *userController) GetUserByID(context *gin.Context) {
	//TODO implement me
	userId := context.Param("user_id")
	newUserId, err := strconv.Atoi(userId)
	if err != nil {
		response.Fail(context, 500, "No User ID found "+err.Error())
		return
	}
	var account, err1 = controller.userService.GetUserByIDService(int(newUserId))
	if err1 != nil {
		response.Fail(context, 500, "not found")
	} else {
		response.Success(context, "successfully", account)
	}

}

func NewUserController(userServ service.UserService) UserController {

	return &userController{userService: userServ}
}
