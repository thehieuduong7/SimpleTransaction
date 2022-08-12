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
            // description: get all my transaction
        }
        GetTransSendedByNumberAcc(ctx *gin.Context){
            // method: Post
            // content_data: MyTransRequest
            // description: get all transaction I sended
        }
        GetTransRevievedByNumberAcc(ctx *gin.Context){
            // method: Post
            // content_data: MyTransRequest
            // description: get all transaction I recieved
        }
    }
