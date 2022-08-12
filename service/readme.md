# Service

## Transaction Service Interface

    type TransactionService interface {
        Create(req dto.CreateTransactionRequest) (dto.CreateTransactionReponse, error){
            // handle insert transaction
        }
        GetAllTransRelatedNumberAcc(dto.GetMyTransactionRequest) ([]dto.GetMyTransactionReponse, error){
            // get my transaction
        }
        GetTransSendedByNumberAcc(dto.GetMyTransactionRequest) ([]dto.GetMyTransactionReponse, error){
            // get transaction I sended
        }
        GetTransRevievedByNumberAcc(dto.GetMyTransactionRequest) ([]dto.GetMyTransactionReponse, error){
            // get transaction I recieved
        }
        GetTransFromTo(dto.GetTransactionFromToRequest) ([]dto.GetMyTransactionReponse, error){
             // get transaction from to
        }

    }
