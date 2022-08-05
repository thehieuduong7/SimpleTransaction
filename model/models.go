package models

import (
	"time"
)

type User struct {
	User_id      int       `json:"user_id"  gorm:"primaryKey;autoIncrement;unique"`
	Name         string    `json:"name" gorm:"NOT NULL"`
	Phone_number string    `json:"phone_number"`
	Email        string    `json:"email"`
	Accounts     []Account `gorm:"foreignKey:User_id;references:User_id;constraint:OnUpdate:CASCADE"`
}

type Account struct {
	Account_number           int           `json:"account_no" gorm:"primaryKey;autoIncrement;unique"`
	User_id                  int           `gorm:"NOT NULL"`
	Surplus                  float64       `json:"surplus" gorm:" check:surplus >= 0"`
	Create_at                time.Time     `json:"create_at" gorm:"autoCreateTime:true"`
	Is_active                bool          `json:"is_active" gorm:"type:bool;default:true"`
	Transaction_resources    []Transaction `gorm:"foreignKey:Account_no_rsc;references:Account_number;constraint:OnUpdate:CASCADE"`
	Transaction_destinations []Transaction `gorm:"foreignKey:Account_no_des;references:Account_number;constraint:OnUpdate:CASCADE"`
}

type Transaction struct {
	Transaction_id int       `json:"transaction_id" gorm:"primaryKey;autoIncrement"`
	Account_no_rsc int       `json:"acc_no_rsc" gorm:"NOT NULL"`
	Account_no_des int       `json:"acc_no_des" gorm:"NOT NULL"`
	Message        string    `json:"message" gorm:"size:500"`
	Amount         float64   `json:"amount" gorm:" check:amount >= 0; NOT NULL"`
	Date_start     time.Time `json:"date_start" `
	Date_success   time.Time `json:"date_success"`
	Is_active      bool      `json:"is_active" gorm:"type:bool;default:true"`
}
