# Respone

## Function

    func Response(ctx *gin.Context, httpStatus int, message string, data interface{}) {
        // base respone
    }

    func Success(ctx *gin.Context, message string, data interface{}) {
        // respone success with ping 200
    }

    func Fail(ctx *gin.Context, httpStatus int, message string) {
        // respone success with ping
    }
