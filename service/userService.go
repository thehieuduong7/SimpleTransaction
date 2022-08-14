package service

import (
	"internBE.com/dto"
	models "internBE.com/entity"
)

type UserService interface {
	CreateUsersService(user dto.UserCreateDTO) models.User
	UpdateUsersService(user dto.UserUpdateDTO)
	DeleteUserService(userId int)
	GetAllUsersService() []dto.UserDTO
	GetUserByIDService(userId int) (dto.UserDTO, error)
}
