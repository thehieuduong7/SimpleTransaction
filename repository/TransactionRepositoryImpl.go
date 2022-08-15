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

func (t *transactionRepositoryImpl) GetHistoryAccountNo(AccountNo int, limit int) (arrAcc []int, err error) {
	result := t.DB.Raw("? union ?",
		t.DB.Model(&entity.Transaction{}).
			Where(entity.Transaction{AccountNoRsc: AccountNo}).
			Distinct("account_no_des"),
		t.DB.Model(&entity.Transaction{}).
			Where(entity.Transaction{AccountNoDes: AccountNo}).
			Distinct("account_no_rsc"),
	).Scan(&arrAcc)

	if result.Error != nil {
		return nil, err
	}

	return arrAcc, nil
}
func (t *transactionRepositoryImpl) GetTransRevievedByNumberAcc(AccountNo int, limit int) (arrTrans []entity.Transaction, err error) {

	result := t.DB.Where(entity.Transaction{AccountNoDes: AccountNo}).Order("create_at desc").Limit(limit).Find(&arrTrans)
	if result.Error != nil {
		return nil, result.Error
	}
	return arrTrans, nil
}

func (t *transactionRepositoryImpl) GetHistoryTransWith(AccountNo1 int,
	AccountNo2 int, limit int) (arrTrans []entity.Transaction, err error) {

	result := t.DB.Where(entity.Transaction{AccountNoRsc: AccountNo1, AccountNoDes: AccountNo2}).
		Or(entity.Transaction{AccountNoRsc: AccountNo2, AccountNoDes: AccountNo1}).
		Order("create_at desc").Limit(limit).Find(&arrTrans)
	if result.Error != nil {
		return nil, result.Error
	}
	return arrTrans, nil
}

func (t *transactionRepositoryImpl) GetHistoryTransBetweenDateWith(AccountNo1 int,
	AccountNo2 int, dateStart time.Time, dateEnd time.Time,
	limit int) (arrTrans []entity.Transaction, err error) {
	result := t.DB.Where(entity.Transaction{AccountNoRsc: AccountNo1, AccountNoDes: AccountNo2}).
		Or(entity.Transaction{AccountNoRsc: AccountNo2, AccountNoDes: AccountNo1}).
		Where("create_at between ? and ?", dateStart, dateEnd).
		Order("create_at desc").Limit(limit).Find(&arrTrans)
	if result.Error != nil {
		return nil, result.Error
	}
	return arrTrans, nil
}

func (t *transactionRepositoryImpl) GetUserByAccountNo(accountNo int) (*entity.User, error) {
	account := entity.Account{}
	if err := t.DB.First(&account, accountNo).Error; err != nil {
		return nil, err
	}
	user := entity.User{}
	if err := t.DB.First(&user, account.UserId).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
