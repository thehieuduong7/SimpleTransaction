package entity

import (
	"time"
)

type User struct {
	UserId      int    `json:"user_id" gorm:"column:user_id;primaryKey;autoIncrement;unique"`
	Name        string `gorm:"NOT NULL"`
	PhoneNumber string
	Email       string
	Accounts    []Account `gorm:"foreignKey:UserId;references:UserId;constraint:OnUpdate:CASCADE"`
}

type Account struct {
	AccountNumber           int           `json:"account_no" gorm:"primaryKey;autoIncrement;unique"`
	UserId                  int           `gorm:"NOT NULL" json:"user_id"`
	Surplus                 float64       `json:"surplus" gorm:" check:surplus >= 0" gorm:"type:decimal(20,8);"`
	CreateAt                time.Time     `json:"create_at" gorm:"autoCreateTime:true"`
	IsActive                bool          `json:"is_active" gorm:"type:bool;default:true"`
	TransactionResources    []Transaction `gorm:"foreignKey:AccountNoRsc;references:AccountNumber;constraint:OnUpdate:CASCADE"`
	TransactionDestinations []Transaction `gorm:"foreignKey:AccountNoDes;references:AccountNumber;constraint:OnUpdate:CASCADE"`
}

type Transaction struct {
	TransactionId int       `gorm:"column:transaction_id;primaryKey;autoIncrement"`
	AccountNoRsc  int       `gorm:"NOT NULL; check:account_no_rsc <> account_no_des"`
	AccountNoDes  int       `gorm:"NOT NULL"`
	Message       string    `gorm:"size:500" `
	Amount        float64   `gorm:"check:amount > 0;NOT NULL" gorm:"type:decimal(20,8);"`
	CreateAt      time.Time `gorm:"autoCreateTime:true"`
}
