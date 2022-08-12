package service

import (
	"internBE.com/dto"
	"internBE.com/entity"
	"internBE.com/repository"
)

type transactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
}

func NewTransactionServiceImpl(transactionRepository *repository.TransactionRepository) TransactionService {
	return &transactionServiceImpl{TransactionRepository: *transactionRepository}
}

func (t *transactionServiceImpl) Create(req dto.CreateTransactionRequest) (dto.CreateTransactionReponse, error) {
	trans := entity.Transaction{
		AccountNoRsc: req.AccountNoRsc,
		AccountNoDes: req.AccountNoDes,
		Message:      req.Message,
		Amount:       req.Amount,
	}
	if err := t.TransactionRepository.Create(&trans); err != nil {
		return dto.CreateTransactionReponse{}, err
	}
	return dto.CreateTransactionReponse{
		AccountNoRsc: trans.AccountNoRsc,
		AccountNoDes: trans.AccountNoDes,
		Message:      trans.Message,
		Amount:       trans.Amount,
		CreateAt:     trans.CreateAt,
	}, nil
}
func (t *transactionServiceImpl) GetAllTransRelatedNumberAcc(req dto.GetMyTransactionRequest) (responses []dto.GetMyTransactionReponse, err error) {
	listTrans, err := t.TransactionRepository.GetAllTransRelatedNumberAcc(req.AccountNoRsc, req.Limit)
	if err != nil {
		return nil, err
	}
	for _, trans := range listTrans {
		responses = append(responses, dto.GetMyTransactionReponse{
			AccountNoRsc: trans.AccountNoRsc,
			AccountNoDes: trans.AccountNoDes,
			Message:      trans.Message,
			Amount:       trans.Amount,
			CreateAt:     trans.CreateAt,
		})
	}
	return responses, err

}
func (t *transactionServiceImpl) GetTransSendedByNumberAcc(req dto.GetMyTransactionRequest) (responses []dto.GetMyTransactionReponse, err error) {
	listTrans, err := t.TransactionRepository.GetTransSendedByNumberAcc(req.AccountNoRsc, req.Limit)
	if err != nil {
		return nil, err
	}
	for _, trans := range listTrans {
		responses = append(responses, dto.GetMyTransactionReponse{
			AccountNoRsc: trans.AccountNoRsc,
			AccountNoDes: trans.AccountNoDes,
			Message:      trans.Message,
			Amount:       trans.Amount,
			CreateAt:     trans.CreateAt,
		})
	}
	return responses, err
}
func (t *transactionServiceImpl) GetTransRevievedByNumberAcc(req dto.GetMyTransactionRequest) (responses []dto.GetMyTransactionReponse, err error) {
	listTrans, err := t.TransactionRepository.GetTransRevievedByNumberAcc(req.AccountNoRsc, req.Limit)
	if err != nil {
		return nil, err
	}
	for _, trans := range listTrans {
		responses = append(responses, dto.GetMyTransactionReponse{
			AccountNoRsc: trans.AccountNoRsc,
			AccountNoDes: trans.AccountNoDes,
			Message:      trans.Message,
			Amount:       trans.Amount,
			CreateAt:     trans.CreateAt,
		})
	}
	return responses, err
}
