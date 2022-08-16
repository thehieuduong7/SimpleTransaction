package repository

import models "internBE.com/entity"

type AccountRepository interface {
	CreateAccount(account *models.Account) *models.Account
	UpdateAccount(account *models.Account) models.Account
	DeleteAccount(id int)
	GetAccountByUserId(id int) []models.Account
	GetAccountById(id int) models.Account
	IsExistAccount(AccountNo int) bool
}
