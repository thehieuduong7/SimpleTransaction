package service

import (
	"log"

	"github.com/mashingan/smapping"
	"internBE.com/dto"
	models "internBE.com/entity"
	"internBE.com/repository"
)

type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *userService {
	return &userService{
		UserRepository: userRepo,
	}
}

func (service *userService) CreateUsersService(user dto.UserCreateDTO) models.User {
	users := models.User{}
	err := smapping.FillStruct(&users, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.UserRepository.CreateUsers(users)
	return res
}

func (service *userService) UpdateUsersService(user dto.UserUpdateDTO) models.User {
	userToUpdate := models.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.UserRepository.UpdateUsers(userToUpdate)
	return updatedUser
}

func (service *userService) DeleteUserService(userId int) {
	//user := models.User{}
	service.UserRepository.DeleteUser(userId)
}

func (service *userService) GetAllUsersService() []models.User {
	return service.UserRepository.GetAllUsers()
}
func (service *userService) GetUserByIDService(userId int) models.User {
	return service.UserRepository.FindUserByID(userId)
}
