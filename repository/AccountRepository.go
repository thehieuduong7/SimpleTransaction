package repository

import (
	"gorm.io/gorm"
	models "internBE.com/model"
)

type AccountRepository interface {
	CreateAccount(account *models.Account) *models.Account
	UpdateAccount(account *models.Account) models.Account
	DeleteAccount(id int)
	GetAccountByUserId(id int) []models.Account
	GetAccountById(id int) models.Account
}
type accountConnection struct {
	connection *gorm.DB
}

func NewAccountRepository(dbConn *gorm.DB) AccountRepository {
	return &accountConnection{
		connection: dbConn,
	}
}

func (db *accountConnection) UpdateAccount(account *models.Account) models.Account {
	var account1 models.Account
	db.connection.Model(models.Account{}).Where("account_number = ?", account.AccountNumber).Updates(&account)
	return account1
}

func (db *accountConnection) DeleteAccount(id int) {
	db.connection.Model(models.Account{}).Where("account_number = ?", id).Update("is_active", false)
}

func (db *accountConnection) GetAccountByUserId(id int) []models.Account {
	var account []models.Account
	db.connection.Find(&account, "user_id = ?", id)
	return account
}

func (db *accountConnection) GetAccountById(id int) models.Account {
	var account models.Account
	db.connection.Find(&account, "account_number = ?", id)
	return account
}

func (db *accountConnection) CreateAccount(account *models.Account) *models.Account {
	db.connection.Create(account)
	return account

}
