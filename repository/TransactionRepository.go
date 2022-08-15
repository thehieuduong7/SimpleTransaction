package repository

import (
	"time"

	"internBE.com/entity"
)

//TransactionRepository interface
type TransactionRepository interface {
	Create(trans *entity.Transaction) error
	GetAllTransRelatedNumberAcc(AccountNo int, limit int) ([]entity.Transaction, error)
	GetHistoryAccountNo(AccountNo int,
		limit int) ([]int, error)
	GetTransRevievedByNumberAcc(AccountNo int,
		limit int) ([]entity.Transaction, error)
	GetHistoryTransWith(AccountNo1 int,
		AccountNo2 int, limit int) ([]entity.Transaction, error)
	GetHistoryTransBetweenDateWith(AccountNo1 int,
		AccountNo2 int, dateStart time.Time, dateEnd time.Time,
		limit int) ([]entity.Transaction, error)
	GetUserByAccountNo(accountNo int) (*entity.User, error)
}
