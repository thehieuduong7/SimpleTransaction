package models

import (
	"time"
)

type User struct {
	User_id      int    `json:"user_id" gorm:"primaryKey;autoIncrement"`
	Name         string `json:"name gorm:NOT NULL"`
	Phone_number string `json:"phone_number"`
	Email        string `json:"email"`
}

type Account struct {
	Account_number int `json:"account_id" gorm:"primaryKey;autoIncrement"`
	User_id        int
	Surplus        float64   `json:"surplus" gorm:" check:surplus >= 0"`
	Create_at      time.Time `json:"create_at" gorm:"autoCreateTime:true"`
	Is_active      bool      `json:"is_active" gorm:"type:bool;default:true"`
}

type Transaction struct {
	Transaction_id int `json:"transaction_id" gorm:"primaryKey;autoIncrement"`
	Account_no_rsc int
	Account_no_des int
	Message        string    `json:"message" gorm:"size:500"`
	Amount         float64   `json:"amount" gorm:" check:amount >= 0; NOT NULL"`
	Date_start     time.Time `json:"date_start" `
	Date_success   time.Time `json:"date_success"`
	Is_success     bool
}
