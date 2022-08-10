package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	models "internBE.com/model"
	"internBE.com/storage"
)

type transactionRepositoryImpl struct {
	DB *gorm.DB
}

func NewTransactionRepositoryImpl() TransactionRepository {
	return &transactionRepositoryImpl{DB: storage.GetDB()}
}
func (t *transactionRepositoryImpl) isAccountActive(AccNo int) (bool, error) {
	user := models.Transaction{}
	result := t.DB.First(&user, AccNo)
	if result.Error != nil {
		return false, result.Error
	}
	return user.IsActive, nil
}

func (t *transactionRepositoryImpl) Create(trans models.Transaction) (*models.Transaction, error) {
	tx := t.DB.Begin()

	//minus money sender
	sender := models.Account{}
	if err := tx.First(&sender, trans.AccountNoRsc).Error; err != nil {
		tx.Rollback()
		return nil, err
	} else if !sender.IsActive {
		tx.Rollback()
		return nil, errors.New("sender not active")
	}
	sender.Surplus -= trans.Amount
	if err := tx.Save(&sender).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// plus money for
	reciver := models.Account{}
	if err := tx.First(&reciver, trans.AccountNoDes).Error; err != nil {
		tx.Rollback()
		return nil, err
	} else if !reciver.IsActive {
		tx.Rollback()
		return nil, errors.New("reciver not active")
	}
	reciver.Surplus += trans.Amount
	if err := tx.Save(&reciver).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// update to history transaction
	if err := tx.Create(&trans).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &trans, nil
}

func (t *transactionRepositoryImpl) GetAllTransRelatedNumberAcc(AccountNo int, limit int) (*[]models.Transaction, error) {
	if checkActive, err := t.isAccountActive(AccountNo); err != nil {
		return nil, err
	} else if !checkActive {
		msg := fmt.Sprintf("Account no %d is not active", AccountNo)
		return nil, errors.New(msg)
	}

	var trans []models.Transaction
	result := t.DB.Where(models.Transaction{AccountNoRsc: AccountNo}).
		Or(models.Transaction{AccountNoDes: AccountNo}).
		Order("transaction_id desc").
		Limit(limit).Find(&trans)
	if result.Error != nil {
		return nil, result.Error
	}
	return &trans, nil
}

// list trans send to
func (t *transactionRepositoryImpl) GetTransSendedByNumberAcc(AccountNo int, limit int) (*[]models.Transaction, error) {
	if checkActive, err := t.isAccountActive(AccountNo); err != nil {
		return nil, err
	} else if !checkActive {
		msg := fmt.Sprintf("Account no %d is not active", AccountNo)
		return nil, errors.New(msg)
	}

	var trans []models.Transaction
	//t.DB.Where("AccountNoRsc").Or(models.Transaction{AccountNoDes: AccountNo}).Find(&trans)

	result := t.DB.Where(models.Transaction{AccountNoRsc: AccountNo}).Order("transaction_id desc").Limit(limit).Find(&trans)
	if result.Error != nil {
		return nil, result.Error
	}
	return &trans, nil
}
func (t *transactionRepositoryImpl) GetTransRevievedByNumberAcc(AccountNo int, limit int) (*[]models.Transaction, error) {
	if checkActive, err := t.isAccountActive(AccountNo); err != nil {
		return nil, err
	} else if !checkActive {
		msg := fmt.Sprintf("Account no %d is not active", AccountNo)
		return nil, errors.New(msg)
	}
	var trans []models.Transaction
	//t.DB.Where("AccountNoRsc").Or(models.Transaction{AccountNoDes: AccountNo}).Find(&trans)

	result := t.DB.Where(models.Transaction{AccountNoDes: AccountNo}).Order("transaction_id desc").Limit(limit).Find(&trans)
	if result.Error != nil {
		return nil, result.Error
	}
	return &trans, nil
}
