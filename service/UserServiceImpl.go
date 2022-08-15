package service

import (
	"internBE.com/storage"
	"log"

	"github.com/mashingan/smapping"
	"internBE.com/dto"
	models "internBE.com/entity"

	"internBE.com/repository"
)

var (
	repoAccount repository.AccountRepository = repository.NewAccountRepository(storage.GetDB())
	serviceRepo AccountService               = NewAccountService(repoAccount)
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

func (service *userService) UpdateUsersService(user dto.UserUpdateDTO) {
	users := models.User{}
	err := smapping.FillStruct(&users, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	service.UserRepository.UpdateUsers(users)

}

func (service *userService) DeleteUserService(userId int) {
	//user := models.User{}
	service.UserRepository.DeleteUser(userId)
}

func (service *userService) GetAllUsersService() []dto.UserDTO {
	users := service.UserRepository.GetAllUsers()
	var lenUsers = len(users)
	var userDtos []dto.UserDTO
	var userDto dto.UserDTO

	for i := 0; i < lenUsers; i++ {
		err := smapping.FillStruct(&userDto, smapping.MapFields(&users[i]))
		if err != nil {
			log.Fatalf("Failed map %v: ", err)
		}
		userDtos = append(userDtos, userDto)
	}
	for i := 0; i < len(userDtos); i++ {
		accounts, _ := serviceRepo.GetAccountByUserId(userDtos[i].UserId)
		userDtos[i].SetAccount(accounts)
	}
	return userDtos
}

// func (service *userService) GetUserByIDService(userId int) models.User {
// 	return service.UserRepository.FindUserByID(userId)
// }

func (service *userService) GetUserByIDService(userId int) (dto.UserDTO, error) {
	user := service.UserRepository.FindUserByID(userId)
	var userDto dto.UserDTO
	err := smapping.FillStruct(&userDto, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	accounts, err := serviceRepo.GetAccountByUserId(userId)
	if err != nil {
		return userDto, nil
	}
	userDto.SetAccount(accounts)
	var err1 = userDto.CheckExisted()
	return userDto, err1
}
