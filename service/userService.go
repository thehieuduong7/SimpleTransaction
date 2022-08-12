package service

import (
	models "internBE.com/model"
	"internBE.com/repository"
)

type UserService interface {
	CreateUsersService(user models.User) models.User
	UpdateUsersService(user models.User) models.User
	DeleteUserByIdService(userId uint) models.User
	GetAllUsersService() []models.User
	FindUserByIDService(userId uint) models.User
}

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: userRepo}
}
