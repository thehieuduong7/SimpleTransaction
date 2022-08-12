package storage

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"internBE.com/entity"
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
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	Host := os.Getenv("DBHost")
	Port := os.Getenv("DBPort")
	User := os.Getenv("DBUser")
	Password := os.Getenv("DBPassword")
	DBName := os.Getenv("DBName")

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
