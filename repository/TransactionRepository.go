package repository

import (
	"internBE.com/entity"
)

//TransactionRepository interface
type TransactionRepository interface {
	Create(trans *entity.Transaction) error
	GetAllTransRelatedNumberAcc(AccountNo int, limit int) ([]entity.Transaction, error)
	GetTransSendedByNumberAcc(AccountNo int, limit int) ([]entity.Transaction, error)
	GetTransRevievedByNumberAcc(AccountNo int, limit int) ([]entity.Transaction, error)
}
