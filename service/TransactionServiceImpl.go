package service

import (
	"errors"
	"time"

	"internBE.com/constantGlobal"
	"internBE.com/dto"
	"internBE.com/entity"
	"internBE.com/repository"
	"internBE.com/storage"
)

type transactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
}

func NewTransactionServiceImpl(transactionRepository *repository.TransactionRepository) TransactionService {
	return &transactionServiceImpl{TransactionRepository: *transactionRepository}
}
func convertCreateEntityToCreateResponeModel(trans entity.Transaction) dto.GetMyTransactionReponse {
	userRepo := repository.NewUserRepository(storage.GetDB())
	Rsc := dto.AccountInfo{Name: userRepo.FindUserByID(trans.AccountNoRsc).Name, AccountNo: trans.AccountNoRsc}
	Des := dto.AccountInfo{Name: userRepo.FindUserByID(trans.AccountNoDes).Name, AccountNo: trans.AccountNoDes}

	return dto.GetMyTransactionReponse{
		Rersouce:    Rsc,
		Destination: Des,
		Message:     trans.Message,
		Amount:      trans.Amount,
		CreateAt:    trans.CreateAt.Format(constantGlobal.TimeLayout),
	}
}

func (t *transactionServiceImpl) Create(req dto.CreateTransactionRequest) (dto.GetMyTransactionReponse, error) {
	trans := entity.Transaction{
		AccountNoRsc: req.AccountNoRsc,
		AccountNoDes: req.AccountNoDes,
		Message:      req.Message,
		Amount:       req.Amount,
	}
	if err := t.TransactionRepository.Create(&trans); err != nil {
		return dto.GetMyTransactionReponse{}, err
	}
	return convertCreateEntityToCreateResponeModel(trans), nil
}
func (t *transactionServiceImpl) GetAllTransRelatedNumberAcc(req dto.GetMyTransactionRequest) (responses []dto.GetMyTransactionReponse, err error) {
	listTrans, err := t.TransactionRepository.GetAllTransRelatedNumberAcc(req.AccountNoRsc, req.Limit)
	if err != nil {
		return nil, err
	}
	for _, trans := range listTrans {
		responses = append(responses, convertCreateEntityToCreateResponeModel(trans))
	}
	return responses, err

}
func (t *transactionServiceImpl) GetTransSendedByNumberAcc(req dto.GetMyTransactionRequest) (responses []dto.GetMyTransactionReponse, err error) {
	listTrans, err := t.TransactionRepository.GetTransSendedByNumberAcc(req.AccountNoRsc, req.Limit)
	if err != nil {
		return nil, err
	}
	for _, trans := range listTrans {
		responses = append(responses, convertCreateEntityToCreateResponeModel(trans))
	}
	return responses, nil
}
func (t *transactionServiceImpl) GetTransRevievedByNumberAcc(req dto.GetMyTransactionRequest) (responses []dto.GetMyTransactionReponse, err error) {
	listTrans, err := t.TransactionRepository.GetTransRevievedByNumberAcc(req.AccountNoRsc, req.Limit)
	if err != nil {
		return nil, err
	}
	for _, trans := range listTrans {
		responses = append(responses, convertCreateEntityToCreateResponeModel(trans))
	}
	return responses, nil
}

func (t *transactionServiceImpl) GetTransFromTo(req dto.GetTransactionFromToRequest) (responses []dto.GetMyTransactionReponse, err error) {
	listTrans, err := t.TransactionRepository.GetTransFromTo(req.AccountNoRsc, req.AccountNoDes, req.Limit)
	if err != nil {
		return nil, err
	}
	for _, trans := range listTrans {
		responses = append(responses, convertCreateEntityToCreateResponeModel(trans))
	}
	return responses, nil
}

func (t *transactionServiceImpl) GetTransDateToDate(req dto.GetTransactionFromToTimeRequest) (responses []dto.
	GetMyTransactionReponse, err error) {
	timeStart, err := time.Parse(constantGlobal.TimeLayout, req.TimeStart)
	if err != nil {
		return nil, err
	}
	timeEnd, err := time.Parse(constantGlobal.TimeLayout, req.TimeEnd)
	if err != nil {
		return nil, err
	}
	if timeStart.After(timeEnd) {
		return nil, errors.New("time start after time end")
	}
	listTrans, err := t.TransactionRepository.
		GetTransDateToDate(req.AccountNoRsc, req.AccountNoDes, timeStart, timeEnd, req.Limit)
	if err != nil {
		return nil, err
	}
	for _, trans := range listTrans {
		responses = append(responses, convertCreateEntityToCreateResponeModel(trans))
	}
	return responses, nil
}
