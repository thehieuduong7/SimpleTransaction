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
	transRouter.POST("/GetHistoryAccountNo", transController.GetHistoryAccountNo)
	transRouter.POST("/TransFromTo", transController.GetTransFromTo)
	transRouter.POST("/GetTransDateToDate", transController.GetTransDateToDate)
	transRouter.POST("/GetStaticTransaction", transController.GetStaticTransaction)

	return r
}
