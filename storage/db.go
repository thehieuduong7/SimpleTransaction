package storage

import (
	"fmt"
	"log"

	models "internBE.com/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
}

var DB *gorm.DB

func Connect(config *Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s  password=%s dbname=%s sslmode=disable",
		config.Host, config.Port,
		config.User,
		config.Password,
		config.DBName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)}) // Log câu lệnh sql trong console

	if err != nil {
		log.Panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Account{}, &models.Transaction{})

	return db
}
