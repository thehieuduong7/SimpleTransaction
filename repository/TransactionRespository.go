package repository

import models "internBE.com/model"

//TransactionRepository interface
type TransactionRepository interface {
	Create(trans models.Transaction) (*models.Transaction, error)
	GetAllTransRelatedNumberAcc(AccountNo int, limit int) (*[]models.Transaction, error)
	GetTransSendedByNumberAcc(AccountNo int, limit int) (*[]models.Transaction, error)
	GetTransRevievedByNumberAcc(AccountNo int, limit int) (*[]models.Transaction, error)
}
