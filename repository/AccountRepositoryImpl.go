package repository

import (
	"gorm.io/gorm"
	models "internBE.com/entity"
)

type accountConnection struct {
	connection *gorm.DB
}

func NewAccountRepository(dbConn *gorm.DB) AccountRepository {
	return &accountConnection{
		connection: dbConn,
	}
}

func (db *accountConnection) UpdateAccount(account *models.Account) (models.Account, error) {
	var account1 models.Account
	err := db.connection.Model(models.Account{}).Where("account_number = ?", account.AccountNumber).Updates(&account).Error
	if err != nil {
		return account1, err
	}
	return account1, nil
}

func (db *accountConnection) DeleteAccount(id int) error {
	err := db.connection.Model(models.Account{}).Where("account_number = ?", id).Update("is_active", false).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *accountConnection) GetAccountByUserId(id int) []models.Account {
	var account []models.Account
	db.connection.Find(&account, "user_id = ? AND is_active= ? ", id, true)
	return account
}

func (db *accountConnection) GetAccountById(id int) models.Account {
	var account models.Account
	db.connection.Find(&account, "account_number = ? AND is_active= ? ", id, true)
	return account
}

func (db *accountConnection) CreateAccount(account *models.Account) (*models.Account, error) {
	err := db.connection.Create(account).Error
	if err != nil {
		return account, err
	}
	return account, nil

}
func (db *accountConnection) IsExistAccount(AccountNo int) bool {
	var account models.Account
	err := db.connection.FindFirst(&account, "account_no = ?", AccountNo).Error
	return err == nil
}
