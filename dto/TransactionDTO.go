package dto

import "time"

type CreateTransactionRequest struct {
	AccountNoRsc int     `json:"account_no_rsc" binding:"required"`
	AccountNoDes int     `json:"account_no_des" binding:"required"`
	Message      string  `json:"message" binding:"required"`
	Amount       float64 `json:"amount" binding:"required"`
}
type CreateTransactionReponse struct {
	AccountNoRsc int       `json:"account_no_rsc"`
	AccountNoDes int       `json:"account_no_des"`
	Message      string    `json:"message"`
	Amount       float64   `json:"amount"`
	CreateAt     time.Time `json:"create_at"`
}
type GetMyTransactionRequest struct {
	AccountNoRsc int `json:"account_no_rsc"`
	Limit        int `json:"limit"`
}

type GetTransactionFromToRequest struct {
	AccountNoRsc int `json:"account_no_rsc"`
	AccountNoDes int `json:"account_no_des"`
	Limit        int `json:"limit"`
}

type GetMyTransactionReponse struct {
	AccountNoRsc int       `json:"account_no_rsc"`
	AccountNoDes int       `json:"account_no_des"`
	Message      string    `json:"message"`
	Amount       float64   `json:"amount"`
	CreateAt     time.Time `json:"create_at"`
}
