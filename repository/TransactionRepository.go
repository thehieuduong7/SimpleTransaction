package repository

import (
	"time"

	"internBE.com/entity"
)

//TransactionRepository interface
type TransactionRepository interface {
	Create(trans *entity.Transaction) error
	GetAllTransRelatedNumberAcc(AccountNo int, limit int) ([]entity.Transaction, error)
	GetTransSendedByNumberAcc(AccountNo int, limit int) ([]entity.Transaction, error)
	GetTransRevievedByNumberAcc(AccountNo int, limit int) ([]entity.Transaction, error)
	GetTransFromTo(AccountNoRsc int, AccountNoDes int, limit int) ([]entity.Transaction, error)
	GetTransDateToDate(AccountNoRsc int, AccountNoDes int, dateStart time.Time, dateEnd time.Time, limit int) ([]entity.Transaction, error)
}
