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
	err := db.connection.Model(models.Account{}).Where("account_no = ?", account.AccountNumber).Updates(&account).Error
	if err != nil {
		return account1, err
	}
	return account1, nil
}

func (db *accountConnection) DeleteAccount(id int) error {
	err := db.connection.Model(models.Account{}).Where("account_no = ?", id).Update("is_active", false).Error
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
	db.connection.Find(&account, "account_no = ? AND is_active= ? ", id, true)
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
	err := db.connection.First(&account, "account_no = ?", AccountNo).Error
	return err == nil
}

func (db *accountConnection) GetUserByAccountNo(accountNo int) (*models.User, error) {
	account := models.Account{}
	if err := db.connection.First(&account, accountNo).Error; err != nil {
		return nil, err
	}
	user := models.User{}
	if err := db.connection.First(&user, account.UserId).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
