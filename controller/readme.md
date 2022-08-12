# Controller

## Transaction Controller Interfaces

    type TransactionController interface {
        Create(ctx *gin.Context){
            // method: Post
            // content_data: CreateRequest
            // description: insert transaction
        }
        GetAllTransRelatedNumberAcc(ctx *gin.Context){
            // method: Post
            // content_data: MyTransRequest
            // description: get my transaction
        }
        GetTransSendedByNumberAcc(ctx *gin.Context){
            // method: Post
            // content_data: MyTransRequest
            // description: get transaction I sended
        }
        GetTransRevievedByNumberAcc(ctx *gin.Context){
            // method: Post
            // content_data: MyTransRequest
            // description: get transaction I recieved
        }
        GetTransFromTo(ctx *gin.Context){
            // method: Post
            // content_data: MyTransFromToRequest
            // description: get transaction from to
        }
    }
