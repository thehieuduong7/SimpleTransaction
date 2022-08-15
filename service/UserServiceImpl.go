package service

import (
	"log"

	"internBE.com/storage"

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
	var userDtos []dto.UserDTO
	for i := 0; i < len(users); i++ {
		userDtos = append(userDtos, makeUserDto(users[i]))
	}
	return userDtos
}

func (service *userService) GetUserByIDService(userId int) (dto.UserDTO, error) {
	user := service.UserRepository.FindUserByID(userId)
	userDto := makeUserDto(user)
	var err1 = userDto.CheckExisted()
	return userDto, err1
}
func makeUserDto(user models.User) dto.UserDTO {
	var userDto dto.UserDTO
	var account dto.AccountDto
	var acounts []dto.AccountDto
	err := smapping.FillStruct(&userDto, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	for i := 0; i < len(user.Accounts); i++ {
		err1 := smapping.FillStruct(&account, smapping.MapFields(&user.Accounts[i]))
		if err1 != nil {
			log.Fatalf("Failed map %v: ", err)
		}
		acounts = append(acounts, account)
	}
	userDto.SetAccount(acounts)
	return userDto
}
