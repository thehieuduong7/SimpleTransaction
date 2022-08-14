// pkg/seeds/users.go
package seeds

import (
	"gorm.io/gorm"
	models "internBE.com/entity"
)

func CreateUser(dbConn *gorm.DB, name string, phoneNumber string, email string) error {
	return dbConn.Create(&models.User{Name: name, PhoneNumber: phoneNumber, Email: email}).Error
}
func CreateAccount(dbConn *gorm.DB, userId int, surplus float64) error {

	return dbConn.Create(&models.Account{UserId: userId, Surplus: surplus}).Error
}
func CreatTransaction(dbConn *gorm.DB, rsc int, des int, message string, amount float64) error {
	return dbConn.Create(&models.Transaction{AccountNoRsc: rsc, AccountNoDes: des, Message: message, Amount: amount}).Error
}
