package main

import (
	"fmt"

	models "internBE.com/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	db.AutoMigrate(&models.User{}, &models.Account{}, &models.Transaction{})
	return db
}
func main() {
	db := Init()
	// user := models.User{Name: "Jinzhu", Phone_number: "123423", Email: "19923"}
	// db.Create(&user) // pass pointer of data to Create
	// fmt.Print(user.User_id)
	// fmt.Print(user.Accounts)

	// user2 := models.User{}
	// db.Find(&user2, 2)
	// account := models.Account{}
	//var user models.User
	users := models.User{}
	var accounts []models.Account
	//fmt.Printf("%+v\n", db.Joins(clause.Associations).Find(&user2, 2)) // Print with Variable Name
	fmt.Printf("%+v\n", db.Preload(clause.Associations).Find(&users, 2)) // Print with Variable Name
	fmt.Printf("%+v\n", users)                                           // Print with Variable Name

	fmt.Print(accounts)

	// acc2 := models.Account{User_id: user2.User_id, Surplus: 400}
	// db.Create(&acc2)          // pass pointer of data to Create
	// fmt.Printf("%+v\n", acc2) // Print with Variable Name

	// fmt.Print(account.Account_number)
	// fmt.Print(account.Is_active)

}
