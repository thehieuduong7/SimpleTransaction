package routes

import (
	"github.com/gin-gonic/gin"
	"internBE.com/controller"
	"internBE.com/repository"
	"internBE.com/service"
	"internBE.com/storage"
)

var (
	repoAccount    repository.AccountRepository = repository.NewAccountRepository(storage.GetDB())
	serviceRepo    service.AccountService       = service.NewAccountService(repoAccount)
	controllerRepo controller.AccountController = controller.NewAccountController(serviceRepo)
)

func AccountRoute(router *gin.Engine) {
	accountRoutes := router.Group("api/account")
	{
		accountRoutes.POST("", controllerRepo.CreateAccount)
		accountRoutes.GET("user/:id", controllerRepo.GetAccountByUserId)
		accountRoutes.GET("/:id", controllerRepo.GetAccountById)
		accountRoutes.PUT("", controllerRepo.UpdateAccount)
		accountRoutes.DELETE("/:id", controllerRepo.DeleteAccount)
	}
}
