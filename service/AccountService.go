package service

import "internBE.com/dto"

type AccountService interface {
	CreateAccount(account *dto.AccountDtoToInsert) error
	UpdateAccount(account *dto.AccountDtoToInsert) error
	DeleteAccount(id int) error
	GetAccountByUserId(id int) ([]dto.AccountDto, error)
	GetAccountById(id int) (dto.AccountDto, error)
}
