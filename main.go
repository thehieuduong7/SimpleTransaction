package main

import (
	"fmt"

	models "internBE.com/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:123456@localhost:5433/demo1"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		fmt.Print("loi")
	}

	db.AutoMigrate(&models.Account{}, &models.User{}, &models.Transaction{})
	return db
}
func main() {
	fmt.Print(Init())
	// Init()

	// user := models.User{Name: "Jinzhu", Phone_number: "123423", Email: "19923"}

	// db.Create(&user) // pass pointer of data to Create

	// fmt.Print(user.User_id)

	// account := models.Account{
	// 	User_id: 2, Surplus: -4,
	// }
	// db.Create(&account) // pass pointer of data to Create

	// fmt.Print(account.Account_number)
	// fmt.Print(account.Is_active)

}
