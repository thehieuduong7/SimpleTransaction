// pkg/seeds/users.go
package seeds

import (
	"gorm.io/gorm"
	models "internBE.com/model"
)

func CreateUser(db *gorm.DB, name string, phone_number string, email string) error {
	return db.Create(&models.User{Name: name, Phone_number: phone_number, Email: email}).Error
}

// Account_Number string  `json:"account_number"`
// 	User_ID        int     `json:"user_id"`
// 	Created_Date   string  `json:"created_date"`
// 	Active_Date    bool    `json:"active_date"`

// func CreateAccount(db *gorm.DB, account_number string, user_id int, surplus float64, created_date string, active_date:bool) error {
// 	return db.Create(&users.User{Account_Number: account_number, User_ID: user_id, Surplus:surplus, Created_Date :created_date,Active_Date:active_date}).Error
// }
