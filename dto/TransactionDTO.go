package dto

type CreateTransactionRequest struct {
	AccountNoRsc int     `json:"account_no_rsc" binding:"required"`
	AccountNoDes int     `json:"account_no_des" binding:"required"`
	Message      string  `json:"message" binding:"required"`
	Amount       float64 `json:"amount" binding:"required"`
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

type GetTransactionFromToTimeRequest struct {
	AccountNoRsc int    `json:"account_no_rsc"`
	AccountNoDes int    `json:"account_no_des"`
	TimeStart    string `json:"time_start"`
	TimeEnd      string `json:"time_end"`
	Limit        int    `json:"limit"`
}

type GetMyTransactionReponse struct {
	Rersouce    AccountInfo `json:"Rersouce"`
	Destination AccountInfo `json:"Destination"`
	Message     string      `json:"message"`
	Amount      float64     `json:"amount"`
	CreateAt    string      `json:"create_at"`
}
type AccountInfo struct {
	Name      string `json:"name"`
	AccountNo int    `json:"account_no"`
}
