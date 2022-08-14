package service

import (
	"internBE.com/dto"
	models "internBE.com/entity"
)

type UserService interface {
	CreateUsersService(user dto.UserCreateDTO)
	UpdateUsersService(user dto.UserUpdateDTO)
	DeleteUserService(userId int)
	GetAllUsersService() []models.User
	GetUserByIDService(userId int) models.User
}
