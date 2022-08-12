package entity

import (
	"time"
)

type User struct {
	UserId      int    `gorm:"column:user_id;primaryKey;autoIncrement;unique"`
	Name        string `gorm:"NOT NULL"`
	PhoneNumber string
	Email       string
	Accounts    []Account `gorm:"foreignKey:UserId;references:UserId;constraint:OnUpdate:CASCADE"`
}

type Account struct {
	AccountNumber           int           `gorm:"column:account_no;primaryKey;autoIncrement;unique"`
	UserId                  int           `gorm:"NOT NULL"`
	Surplus                 float64       `gorm:"check:surplus >= 0"`
	CreateAt                time.Time     `gorm:"autoCreateTime:true"`
	IsActive                bool          `gorm:"type:bool;default:true"`
	TransactionResources    []Transaction `gorm:"foreignKey:AccountNoRsc;references:AccountNumber;constraint:OnUpdate:CASCADE"`
	TransactionDestinations []Transaction `gorm:"foreignKey:AccountNoDes;references:AccountNumber;constraint:OnUpdate:CASCADE"`
}

type Transaction struct {
	TransactionId int       `gorm:"column:transaction_id;primaryKey;autoIncrement"`
	AccountNoRsc  int       `gorm:"NOT NULL"`
	AccountNoDes  int       `gorm:"NOT NULL"`
	Message       string    `gorm:"size:500"`
	Amount        float64   `gorm:"check:amount > 0;NOT NULL"`
	CreateAt      time.Time `gorm:"autoCreateTime:true"`
}
