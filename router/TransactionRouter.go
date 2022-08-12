package router

import (
	"github.com/gin-gonic/gin"
	"internBE.com/controller"
)

func CollectTransactionRoute(r *gin.Engine, transController controller.TransactionController) *gin.Engine {
	transRouter := r.Group("/transaction")
	transRouter.POST("", transController.Create)

	transRouter.POST("/AllMyTrans", transController.GetAllTransRelatedNumberAcc)
	transRouter.POST("/AllSendFrom", transController.GetTransSendedByNumberAcc)
	transRouter.POST("/AllRecieveBy", transController.GetTransRevievedByNumberAcc)

	return r
}
