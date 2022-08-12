# Service

## Transaction Service Interface

    type TransactionService interface {
        Create(req dto.CreateTransactionRequest) (dto.CreateTransactionReponse, error){
            // handle insert transaction
        }
        GetAllTransRelatedNumberAcc(dto.GetMyTransactionRequest) ([]dto.GetMyTransactionReponse, error){
            // get all my transaction
        }
        GetTransSendedByNumberAcc(dto.GetMyTransactionRequest) ([]dto.GetMyTransactionReponse, error){
            // get all transaction I sended
        }
        GetTransRevievedByNumberAcc(dto.GetMyTransactionRequest) ([]dto.GetMyTransactionReponse, error){
            // get all transaction I recieved
        }
    }
