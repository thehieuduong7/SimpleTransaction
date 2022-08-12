package dto

import "errors"

type AccountDtoToInsert struct {
	AccountNumber int     `json:"account_no"`
	Surplus       float64 `json:"surplus" binding:"required"`
	UserId        int     `json:"user_id" binding:"required"`
}
type AccountDto struct {
	AccountNumber int     `json:"account_no"`
	Surplus       float64 `json:"surplus"`
}

func (account *AccountDtoToInsert) SetAccountNumber(id int) {
	account.AccountNumber = id
}
func (account *AccountDto) CheckExisted() error {

	if account.AccountNumber == 0 {
		return errors.New("not existed account ")
	}
	return nil
}
