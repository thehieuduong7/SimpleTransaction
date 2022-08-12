# Repository

## TransactionRepository

    type TransactionRepository interface {
        Create(trans *entity.Transaction) error{
            // insert transaction
        }
        GetAllTransRelatedNumberAcc(AccountNo int, limit int) ([]entity.Transaction, error) {
            // get my trasaction
        }
        GetTransSendedByNumberAcc(AccountNo int, limit int) ([]entity.Transaction, error) {
            // get transaction I sended
        }
        GetTransRevievedByNumberAcc(AccountNo int, limit int) ([]entity.Transaction, error) {
            // get transaction I received
        }
        GetTransFromTo(AccountNoRsc int, AccountNoDes int, limit int) ([]entity.Transaction, error){
            // get transaction from to
        }
    }
