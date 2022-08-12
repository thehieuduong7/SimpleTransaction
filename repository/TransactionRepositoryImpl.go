package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	"internBE.com/entity"
	"internBE.com/storage"
)

type transactionRepositoryImpl struct {
	DB *gorm.DB
}

func NewTransactionRepositoryImpl() TransactionRepository {
	return &transactionRepositoryImpl{DB: storage.GetDB()}
}
func (t *transactionRepositoryImpl) isAccountActive(AccNo int) (bool, error) {
	user := entity.Account{}
	result := t.DB.First(&user, AccNo)
	if result.Error != nil {
		return false, result.Error
	}
	return user.IsActive, nil
}

func (t *transactionRepositoryImpl) Create(trans *entity.Transaction) error {
	tx := t.DB.Begin()

	//minus money sender
	sender := entity.Account{}
	if err := tx.First(&sender, trans.AccountNoRsc).Error; err != nil {
		tx.Rollback()
		return err
	} else if !sender.IsActive {
		tx.Rollback()
		return errors.New("sender not active")
	}
	sender.Surplus -= trans.Amount
	if err := tx.Save(&sender).Error; err != nil {
		tx.Rollback()
		return err
	}

	// plus money for
	reciver := entity.Account{}
	if err := tx.First(&reciver, trans.AccountNoDes).Error; err != nil {
		tx.Rollback()
		return err
	} else if !reciver.IsActive {
		tx.Rollback()
		return errors.New("reciver not active")
	}
	reciver.Surplus += trans.Amount
	if err := tx.Save(&reciver).Error; err != nil {
		tx.Rollback()
		return err
	}

	// update to history transaction
	if err := tx.Create(&trans).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (t *transactionRepositoryImpl) GetAllTransRelatedNumberAcc(AccountNo int, limit int) ([]entity.Transaction, error) {
	if checkActive, err := t.isAccountActive(AccountNo); err != nil {
		return nil, err
	} else if !checkActive {
		msg := fmt.Sprintf("Account no %d is not active", AccountNo)
		return nil, errors.New(msg)
	}

	var trans []entity.Transaction
	result := t.DB.Where(entity.Transaction{AccountNoRsc: AccountNo}).
		Or(entity.Transaction{AccountNoDes: AccountNo}).
		Order("transaction_id desc").
		Limit(limit).Find(&trans)
	if result.Error != nil {
		return nil, result.Error
	}
	return trans, nil
}

// list trans send to
func (t *transactionRepositoryImpl) GetTransSendedByNumberAcc(AccountNo int, limit int) ([]entity.Transaction, error) {
	if checkActive, err := t.isAccountActive(AccountNo); err != nil {
		return nil, err
	} else if !checkActive {
		msg := fmt.Sprintf("Account no %d is not active", AccountNo)
		return nil, errors.New(msg)
	}

	var trans []entity.Transaction
	result := t.DB.Where(entity.Transaction{AccountNoRsc: AccountNo}).Order("transaction_id desc").Limit(limit).Find(&trans)
	if result.Error != nil {
		return nil, result.Error
	}
	return trans, nil
}
func (t *transactionRepositoryImpl) GetTransRevievedByNumberAcc(AccountNo int, limit int) ([]entity.Transaction, error) {
	if checkActive, err := t.isAccountActive(AccountNo); err != nil {
		return nil, err
	} else if !checkActive {
		msg := fmt.Sprintf("Account no %d is not active", AccountNo)
		return nil, errors.New(msg)
	}
	var trans []entity.Transaction
	result := t.DB.Where(entity.Transaction{AccountNoDes: AccountNo}).Order("transaction_id desc").Limit(limit).Find(&trans)
	if result.Error != nil {
		return nil, result.Error
	}
	return trans, nil
}
