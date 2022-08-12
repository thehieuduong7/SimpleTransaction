package routes

import (
	"github.com/gin-gonic/gin"
	"internBE.com/controller"
	"internBE.com/repository"
	"internBE.com/service"
	"internBE.com/storage"
)

var (
	transRepository repository.TransactionRepository = repository.NewTransactionRepositoryImpl(storage.GetDB())
	transService    service.TransactionService       = service.NewTransactionServiceImpl(&transRepository)
	transController controller.TransactionController = controller.NewTransactionControllerImpl(&transService)
)

func TransactionRoute(r *gin.Engine) *gin.Engine {
	transRouter := r.Group("/transaction")
	transRouter.POST("", transController.Create)
	transRouter.POST("/MyTrans", transController.GetAllTransRelatedNumberAcc)
	transRouter.POST("/TransSendFrom", transController.GetTransSendedByNumberAcc)
	transRouter.POST("/TransRecieveBy", transController.GetTransRevievedByNumberAcc)
	transRouter.POST("/TransFromTo", transController.GetTransFromTo)
	return r
}
