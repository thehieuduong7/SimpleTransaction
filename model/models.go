package models

import (
	"time"
)

type User struct {
	UserID      int       `json:"user_id"  gorm:"primaryKey;autoIncrement;unique"`
	Name        string    `json:"name" gorm:"NOT NULL"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Accounts    []Account `gorm:"foreignKey:UserID;references:UserID;constraint:OnUpdate:CASCADE"`
}

type Account struct {
	AccountNumber           int           `json:"account_no" gorm:"primaryKey;autoIncrement;unique"`
	UserId                  int           `gorm:"NOT NULL"`
	Surplus                 float64       `json:"surplus" gorm:" check:surplus >= 0"`
	CreateAt                time.Time     `json:"create_at" gorm:"autoCreateTime:true"`
	IsActive                bool          `json:"is_active" gorm:"type:bool;default:true"`
	TransactionResources    []Transaction `gorm:"foreignKey:AccountNoRsc;references:AccountNumber;constraint:OnUpdate:CASCADE"`
	TransactionDestinations []Transaction `gorm:"foreignKey:AccountNoDes;references:AccountNumber;constraint:OnUpdate:CASCADE"`
}

type Transaction struct {
	TransactionId int       `json:"transaction_id" gorm:"primaryKey;autoIncrement"`
	AccountNoRsc  int       `json:"acc_no_rsc" gorm:"NOT NULL"`
	AccountNoDes  int       `json:"acc_no_des" gorm:"NOT NULL"`
	Message       string    `json:"message" gorm:"size:500"`
	Amount        float64   `json:"amount" gorm:" check:amount > 0; NOT NULL"`
	DateStart     time.Time `json:"date_start" `
	DateSuccess   time.Time `json:"date_success"`
	IsActive      bool      `json:"is_active" gorm:"type:bool;default:true"`
}
