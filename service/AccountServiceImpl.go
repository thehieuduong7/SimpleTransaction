package service

import (
	"errors"
	"log"

	"github.com/mashingan/smapping"
	"internBE.com/dto"
	"internBE.com/entity"
	"internBE.com/repository"
)

type accountService struct {
	AccountRepo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) AccountService {
	return &accountService{AccountRepo: repo}
}

func (service *accountService) CreateAccount(account *dto.AccountDtoToInsert) {
	var accountEntity = entity.Account{}
	err := smapping.FillStruct(&accountEntity, smapping.MapFields(account))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	accountEntity2 := service.AccountRepo.CreateAccount(&accountEntity)
	account.SetAccountNumber(accountEntity2.AccountNumber)
}

func (service *accountService) UpdateAccount(account *dto.AccountDtoToInsert) {
	var accountEntity = entity.Account{}
	err := smapping.FillStruct(&accountEntity, smapping.MapFields(account))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	service.AccountRepo.UpdateAccount(&accountEntity)
}

func (service *accountService) DeleteAccount(id int) {
	service.AccountRepo.DeleteAccount(id)
}

func (service *accountService) GetAccountByUserId(id int) ([]dto.AccountDto, error) {
	account := service.AccountRepo.GetAccountByUserId(id)
	var accountDtos []dto.AccountDto
	var accountDto dto.AccountDto
	for _, acc := range account {
		err := smapping.FillStruct(&accountDto, smapping.MapFields(&acc))
		if err != nil {
			log.Fatalf("Failed map %v: ", err)
		}
		accountDtos = append(accountDtos, accountDto)
	}
	var err error = nil
	if len(accountDtos) == 0 {
		err = errors.New("null")
	}
	return accountDtos, err
}

func (service *accountService) GetAccountById(id int) (dto.AccountDto, error) {
	account := service.AccountRepo.GetAccountById(id)
	var accountDto dto.AccountDto
	err := smapping.FillStruct(&accountDto, smapping.MapFields(&account))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	var err1 = accountDto.CheckExisted()
	return accountDto, err1
}
