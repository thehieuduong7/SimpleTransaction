package main

import (
	"fmt"

	models "internBE.com/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:123456@localhost:5433/demo1"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info), // Log câu lệnh sql trong console
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	if err != nil {
		fmt.Print("loi")
	}

	db.Migrator().CreateTable(&models.User{}, &models.Account{}, &models.Transaction{})
	return db
}
func main() {
	Init()
	// user := models.User{Name: "Jinzhu", Phone_number: "123423", Email: "19923"}
	// db.Create(&user) // pass pointer of data to Create
	// fmt.Print(user.User_id)
	// fmt.Print(user.Accounts)

	// account := models.Account{
	// 	User_id: 2, Surplus: -4,
	// }
	// db.Create(&account) // pass pointer of data to Create

	// fmt.Print(account.Account_number)
	// fmt.Print(account.Is_active)

}
