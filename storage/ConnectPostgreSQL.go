package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"internBE.com/constant"
	"internBE.com/entity"
	"log"
)

//type Config struct {
//	Host     string
//	Port     string
//	Password string
//	User     string
//	DBName   string
//}

var DB *gorm.DB

func Connect() *gorm.DB {

	Host := constant.Host
	Port := constant.Port
	User := constant.User
	Password := constant.Password
	DBName := constant.DBName

	dsn := fmt.Sprintf("host=%s port=%s user=%s  password=%s dbname=%s sslmode=disable",
		Host, Port,
		User,
		Password,
		DBName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log câu lệnh sql trong console

	})
	if err != nil {
		log.Panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Account{}, &entity.Transaction{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	if DB == nil {
		return Connect()
	}
	return DB
}
