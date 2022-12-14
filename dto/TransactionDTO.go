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
	AccountNo1 int `json:"account_no_1"`
	AccountNo2 int `json:"account_no_2"`
	Limit      int `json:"limit"`
}

type GetTransactionFromToTimeRequest struct {
	AccountNo1 int    `json:"account_no_1"`
	AccountNo2 int    `json:"account_no_2"`
	TimeStart  string `json:"time_start"`
	TimeEnd    string `json:"time_end"`
	Limit      int    `json:"limit"`
}

type GetMyTransactionReponse struct {
	Resource    AccountInfo `json:"Resource"`
	Destination AccountInfo `json:"Destination"`
	Message     string      `json:"message"`
	Amount      float64     `json:"amount"`
	CreateAt    string      `json:"create_at"`
}
type AccountInfo struct {
	Name      string `json:"name"`
	AccountNo int    `json:"account_no"`
}
type StaticTransactionRequest struct {
	AccountNo1 int `json:"account_no_1"`
	AccountNo2 int `json:"account_no_2"`
}

type StaticTransactionResponse struct {
	Resource       AccountInfo  `json:"resource"`
	Destination    AccountInfo  `json:"destination"`
	StaticSended   StaticDetail `json:"static_sended"`
	StaticRecieved StaticDetail `json:"count_recieved_transaction"`
}

type StaticDetail struct {
	CountTransaction int64   `json:"count_transaction"`
	TotalMoney       float64 `json:"total_money"`
}
