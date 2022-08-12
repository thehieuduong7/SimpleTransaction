package service

import "internBE.com/dto"

type AccountService interface {
	CreateAccount(account *dto.AccountDtoToInsert)
	UpdateAccount(account *dto.AccountDtoToInsert)
	DeleteAccount(id int)
	GetAccountByUserId(id int) ([]dto.AccountDto, error)
	GetAccountById(id int) (dto.AccountDto, error)
}
