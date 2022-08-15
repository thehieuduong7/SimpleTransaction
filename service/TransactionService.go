package service

import (
	"internBE.com/dto"
)

type TransactionService interface {
	Create(req dto.CreateTransactionRequest) (dto.GetMyTransactionReponse, error)
	GetAllTransRelatedNumberAcc(dto.GetMyTransactionRequest) ([]dto.GetMyTransactionReponse, error)
	GetHistoryAccountNo(dto.GetMyTransactionRequest) ([]dto.AccountInfo, error)
	GetTransRevievedByNumberAcc(dto.GetMyTransactionRequest) ([]dto.GetMyTransactionReponse, error)
	GetHistoryTransWith(dto.GetTransactionFromToRequest) ([]dto.GetMyTransactionReponse, error)
	GetHistoryTransBetweenDateWith(dto.GetTransactionFromToTimeRequest) ([]dto.GetMyTransactionReponse, error)
}
