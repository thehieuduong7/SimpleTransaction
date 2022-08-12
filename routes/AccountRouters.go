package routes

import (
	"github.com/gin-gonic/gin"
	"internBE.com/constant"
	"internBE.com/controller"
	"internBE.com/repository"
	"internBE.com/service"
	"internBE.com/storage"
)

var (
	config = &storage.Config{
		Host:     constant.Host,
		Port:     constant.Port,
		Password: constant.Password,
		User:     constant.User,
		DBName:   constant.DBName}
	repoAccount    repository.AccountRepository = repository.NewAccountRepository(storage.Connect(config))
	serviceRepo    service.AccountService       = service.NewAccountService(repoAccount)
	controllerRepo controller.AccountController = controller.NewAccountController(serviceRepo)
)

func AccountRoute(router *gin.Engine) {
	accountRoutes := router.Group("api/account")
	{
		accountRoutes.POST("/create", controllerRepo.CreateAccount)
		accountRoutes.GET("user/:id", controllerRepo.GetAccountByUserId)
		accountRoutes.GET("/:id", controllerRepo.GetAccountById)
		accountRoutes.PUT("/update", controllerRepo.UpdateAccount)
		accountRoutes.DELETE("/:id", controllerRepo.DeleteAccount)
	}
}
