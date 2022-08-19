package service

import (
	"errors"
	"fmt"
	"time"

	"internBE.com/constantGlobal"
	"internBE.com/dto"
	"internBE.com/entity"
	"internBE.com/repository"
)

type transactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
	AccountRepository     repository.AccountRepository
}

func NewTransactionServiceImpl(transactionRepository *repository.TransactionRepository,
	accountRepository *repository.AccountRepository) TransactionService {

	return &transactionServiceImpl{TransactionRepository: *transactionRepository,
		AccountRepository: *accountRepository,
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
	return t.convertCreateEntityToCreateResponeModel(trans), nil
}
func (t *transactionServiceImpl) GetAllTransRelatedNumberAcc(req dto.GetMyTransactionRequest) (responses []dto.GetMyTransactionReponse, err error) {
	if err := t.IsExistAccount(req.AccountNoRsc); err != nil {
		return nil, err
	}
	listTrans, err := t.TransactionRepository.GetAllTransRelatedNumberAcc(req.AccountNoRsc, req.Limit)
	if err != nil {
		return nil, err
	}
	for _, trans := range listTrans {
		responses = append(responses, t.convertCreateEntityToCreateResponeModel(trans))
	}
	return responses, err

}
func (t *transactionServiceImpl) GetHistoryAccountNo(req dto.GetMyTransactionRequest) (responses []dto.AccountInfo, err error) {
	if err := t.IsExistAccount(req.AccountNoRsc); err != nil {
		return nil, err
	}
	list, err := t.TransactionRepository.GetHistoryAccountNo(req.AccountNoRsc, req.Limit)
	if err != nil {
		return nil, err
	}

	for _, account_no := range list {
		user, _ := t.AccountRepository.GetUserByAccountNo(account_no)
		Rsc := dto.AccountInfo{Name: user.Name, AccountNo: account_no}
		responses = append(responses, Rsc)
	}
	return responses, nil
}

func (t *transactionServiceImpl) GetHistoryTransWith(req dto.GetTransactionFromToRequest) (responses []dto.GetMyTransactionReponse, err error) {
	if err := t.IsExistAccount(req.AccountNo1); err != nil {
		return nil, err
	}
	if err := t.IsExistAccount(req.AccountNo2); err != nil {
		return nil, err
	}

	listTrans, err := t.TransactionRepository.
		GetHistoryTransWith(req.AccountNo1, req.AccountNo2, req.Limit)
	if err != nil {
		return nil, err
	}
	for _, trans := range listTrans {
		responses = append(responses, t.convertCreateEntityToCreateResponeModel(trans))
	}
	return responses, nil
}

func (t *transactionServiceImpl) GetHistoryTransBetweenDateWith(req dto.GetTransactionFromToTimeRequest) (responses []dto.
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
	if err := t.IsExistAccount(req.AccountNo1); err != nil {
		return nil, err
	}
	if err := t.IsExistAccount(req.AccountNo2); err != nil {
		return nil, err
	}

	listTrans, err := t.TransactionRepository.
		GetHistoryTransBetweenDateWith(req.AccountNo1, req.AccountNo2, timeStart, timeEnd, req.Limit)
	if err != nil {
		return nil, err
	}
	for _, trans := range listTrans {
		responses = append(responses, t.convertCreateEntityToCreateResponeModel(trans))
	}
	return responses, nil
}
func (t *transactionServiceImpl) IsExistAccount(accountNo int) error {
	if !t.AccountRepository.IsExistAccount(accountNo) {
		msg := fmt.Sprintf("account no %d not found", accountNo)
		return errors.New(msg)
	}
	return nil
}

func (t *transactionServiceImpl) GetStaticTransaction(req dto.
	StaticTransactionRequest) (static dto.StaticTransactionResponse, err error) {
	if err := t.IsExistAccount(req.AccountNo1); err != nil {
		return static, err
	}
	if err := t.IsExistAccount(req.AccountNo2); err != nil {
		return static, err
	}
	static.Resource = t.convertToAccountDetailDTO(req.AccountNo1)
	static.Destination = t.convertToAccountDetailDTO(req.AccountNo2)
	static.StaticSended, _ = t.convertToStaticDetail(req.AccountNo1, req.AccountNo2)
	static.StaticRecieved, _ = t.convertToStaticDetail(req.AccountNo2, req.AccountNo1)
	return static, nil
}

func (t *transactionServiceImpl) convertToAccountDetailDTO(accountNo int) dto.AccountInfo {
	user, _ := t.AccountRepository.GetUserByAccountNo(accountNo)
	return dto.AccountInfo{Name: user.Name, AccountNo: accountNo}
}
func (t *transactionServiceImpl) convertToStaticDetail(AccountNo1 int,
	AccountNo2 int) (dto.StaticDetail, error) {
	count, money, err := t.TransactionRepository.
		GetStaticTransaction(AccountNo1, AccountNo2)

	if err != nil {
		return dto.StaticDetail{}, err
	}
	return dto.StaticDetail{CountTransaction: count, TotalMoney: money}, nil
}
func (t *transactionServiceImpl) convertCreateEntityToCreateResponeModel(
	trans entity.Transaction) dto.GetMyTransactionReponse {
	Rsc := t.convertToAccountDetailDTO(trans.AccountNoRsc)
	Des := t.convertToAccountDetailDTO(trans.AccountNoDes)

	return dto.GetMyTransactionReponse{
		Resource:    Rsc,
		Destination: Des,
		Message:     trans.Message,
		Amount:      trans.Amount,
		CreateAt:    trans.CreateAt.Format(constantGlobal.TimeLayout),
	}
}
