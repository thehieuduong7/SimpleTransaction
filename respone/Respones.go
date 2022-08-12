package response

import (
	"github.com/gin-gonic/gin"
)

//response is used for static shape json return
type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//EmptyObj object is used when data doesnt want to be null on json
type emptyOb struct{}

func Response(ctx *gin.Context, httpStatus int, message string, data interface{}) {
	ctx.JSON(httpStatus, response{
		Status:  httpStatus,
		Message: message,
		Data:    data,
	})
}

//Buildresponse method is to inject data value to dynamic success response
func Success(ctx *gin.Context, message string, data interface{}) {
	Response(ctx, 200, message, data)
}

//BuildErrorresponse method is to inject data value to dynamic failed response
func Fail(ctx *gin.Context, httpStatus int, message string) {
	Response(ctx, httpStatus, message, emptyOb{})
}
