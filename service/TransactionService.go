package service

import (
	"internBE.com/dto"
)

type TransactionService interface {
	Create(req dto.CreateTransactionRequest) (dto.CreateTransactionReponse, error)
	GetAllTransRelatedNumberAcc(dto.GetMyTransactionRequest) ([]dto.GetMyTransactionReponse, error)
	GetTransSendedByNumberAcc(dto.GetMyTransactionRequest) ([]dto.GetMyTransactionReponse, error)
	GetTransRevievedByNumberAcc(dto.GetMyTransactionRequest) ([]dto.GetMyTransactionReponse, error)
}
