package repository

import (
	models "internBE.com/entity"
)

type AccountRepository interface {
	CreateAccount(account *models.Account) (*models.Account, error)
	UpdateAccount(account *models.Account) (models.Account, error)
	DeleteAccount(id int) error
	GetAccountByUserId(id int) []models.Account
	GetAccountById(id int) models.Account
	IsExistAccount(AccountNo int) bool
	GetUserByAccountNo(accountNo int) (*models.User, error)
}
