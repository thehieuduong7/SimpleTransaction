package service

import (
	"internBE.com/dto"
	models "internBE.com/entity"
)

type UserService interface {
	CreateUsersService(user dto.UserCreateDTO) models.User
	UpdateUsersService(user dto.UserUpdateDTO) models.User
	DeleteUserService(userId int)
	GetAllUsersService() []models.User
	GetUserByIDService(userId int) models.User
}
