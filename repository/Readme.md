# Repository

## TransactionRepository

    define interface
    func:

    1. Create(trans model.Transaction) (\*model.Transaction, error)
    // insert into history transaction and update surplus money account
    2. GetAllTransRelatedNumberAcc(AccountNo int, limit int) (_[]model.Transaction, error)
    // get top ? _ from transactions where account_no_rsc = ? or account_no_des = ?
    3. GetTransSendedByNumberAcc(AccountNo int, limit int) (\*[]model.Transaction, error)
    // get top ? from transactions where account_no_rsc = ?
    4. GetTransRevievedByNumberAcc(AccountNo int, limit int) (\*[]model.Transaction, error)
    // get top ? from transactions where account_no_des = ?

## TransactionRepositoryImpl

    define implement
    func:
    more...

    1. func NewTransactionRepositoryImpl() TransactionRepository
    // constructor
    2. func (t \*transactionRepositoryImpl) isAccountActive(AccNo int) (bool, error)
    // private, check account active
