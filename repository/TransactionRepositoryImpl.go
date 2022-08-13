package repository

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"internBE.com/entity"
)

type transactionRepositoryImpl struct {
	DB *gorm.DB
}

func NewTransactionRepositoryImpl(db *gorm.DB) TransactionRepository {
	return &transactionRepositoryImpl{DB: db}
}
func plusMoneyAccount(accountNo int, money float64, tx *gorm.DB) error {
	account := entity.Account{}
	if err := tx.First(&account, accountNo).Error; err != nil {
		return err
	} else if !account.IsActive {
		msg := fmt.Sprintf("Account no %d is not active", accountNo)
		return errors.New(msg)
	}
	account.Surplus += money
	return tx.Save(&account).Error
}

func (t *transactionRepositoryImpl) Create(trans *entity.Transaction) error {
	tx := t.DB.Begin()

	//minus money sender
	if err := plusMoneyAccount(trans.AccountNoDes, -trans.Amount, tx); err != nil {
		tx.Rollback()
		return err
	}

	// plus money for
	if err := plusMoneyAccount(trans.AccountNoRsc, trans.Amount, tx); err != nil {
		tx.Rollback()
		return err
	}

	// update to history transaction
	trans.CreateAt = time.Now()
	tx.Create(&trans)

	//commit
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (t *transactionRepositoryImpl) GetAllTransRelatedNumberAcc(AccountNo int, limit int) (arrTrans []entity.Transaction, err error) {

	result := t.DB.Where(entity.Transaction{AccountNoRsc: AccountNo}).
		Or(entity.Transaction{AccountNoDes: AccountNo}).
		Order("create_at desc").
		Limit(limit).Find(&arrTrans)
	if result.Error != nil {
		return nil, result.Error
	}
	return arrTrans, nil
}

// list trans send to
func (t *transactionRepositoryImpl) GetTransSendedByNumberAcc(AccountNo int, limit int) (arrTrans []entity.Transaction, err error) {

	result := t.DB.Where(entity.Transaction{AccountNoRsc: AccountNo}).Order("create_at desc").Limit(limit).Find(&arrTrans)
	if result.Error != nil {
		return nil, result.Error
	}
	return arrTrans, nil
}
func (t *transactionRepositoryImpl) GetTransRevievedByNumberAcc(AccountNo int, limit int) (arrTrans []entity.Transaction, err error) {

	result := t.DB.Where(entity.Transaction{AccountNoDes: AccountNo}).Order("create_at desc").Limit(limit).Find(&arrTrans)
	if result.Error != nil {
		return nil, result.Error
	}
	return arrTrans, nil
}

func (t *transactionRepositoryImpl) GetTransFromTo(AccountNoRsc int, AccountNoDes int, limit int) (arrTrans []entity.Transaction, err error) {

	result := t.DB.Where(entity.Transaction{AccountNoRsc: AccountNoRsc, AccountNoDes: AccountNoDes}).
		Order("create_at desc").Limit(limit).Find(&arrTrans)
	if result.Error != nil {
		return nil, result.Error
	}
	return arrTrans, nil
}

func (t *transactionRepositoryImpl) GetTransDateToDate(AccountNoRsc int,
	AccountNoDes int, dateStart time.Time, dateEnd time.Time,
	limit int) (arrTrans []entity.Transaction, err error) {
	result := t.DB.Where(entity.Transaction{AccountNoRsc: AccountNoRsc, AccountNoDes: AccountNoDes}).
		Where("create_at between ? and ?", dateStart, dateEnd).
		Order("create_at desc").Limit(limit).Find(&arrTrans)
	if result.Error != nil {
		return nil, result.Error
	}
	return arrTrans, nil
}
