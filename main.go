package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"internBE.com/controller"
	"internBE.com/repository"
	"internBE.com/router"
	"internBE.com/service"
	"internBE.com/storage"
)

func main() {
	err := godotenv.Load(".env")

	config := &storage.Config{
		Host:     os.Getenv("DBHost"),
		Port:     os.Getenv("DBPort"),
		Password: os.Getenv("DBPassword"),
		User:     os.Getenv("DBUser"),
		DBName:   os.Getenv("DBName")}

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	storage.Connect(config)

	res := repository.NewTransactionRepositoryImpl()
	ser := service.NewTransactionServiceImpl(&res)
	con := controller.NewTransactionControllerImpl(&ser)

	r := gin.New()
	// call router
	router.CollectTransactionRoute(r, con)
	r.Run(":8000")
	// trans := entity.Transaction{AccountNoRsc: 1, AccountNoDes: 2, Amount: 10, Message: "cam on 10k"}
	// fmt.Print(res.Create(&trans))
	// t.Log(transRepo.Create(trans))
}
