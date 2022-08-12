package service

import (
	"errors"
	"github.com/mashingan/smapping"
	"internBE.com/dto"
	"internBE.com/entity"
	"internBE.com/repository"
	"log"
)

type accountService struct {
	AccountRepo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) AccountService {
	return &accountService{AccountRepo: repo}
}

func (controller *accountService) CreateAccount(account *dto.AccountDtoToInsert) {
	var accountEntity = entity.Account{}
	err := smapping.FillStruct(&accountEntity, smapping.MapFields(account))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	accountEntity2 := controller.AccountRepo.CreateAccount(&accountEntity)
	account.SetAccountNumber(accountEntity2.AccountNumber)
}

func (controller *accountService) UpdateAccount(account *dto.AccountDtoToInsert) {
	var accountEntity = entity.Account{}
	err := smapping.FillStruct(&accountEntity, smapping.MapFields(account))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	controller.AccountRepo.UpdateAccount(&accountEntity)
}

func (controller *accountService) DeleteAccount(id int) {
	controller.AccountRepo.DeleteAccount(id)
}

func (controller *accountService) GetAccountByUserId(id int) ([]dto.AccountDto, error) {
	account := controller.AccountRepo.GetAccountByUserId(id)
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

func (controller *accountService) GetAccountById(id int) (dto.AccountDto, error) {
	account := controller.AccountRepo.GetAccountById(id)
	var accountDto dto.AccountDto
	err := smapping.FillStruct(&accountDto, smapping.MapFields(&account))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	var err1 = accountDto.CheckExisted()
	return accountDto, err1
}
