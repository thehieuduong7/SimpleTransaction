package repository

import (
	"fmt"
	"testing"

	"internBE.com/entity"
)

func Test_create(t *testing.T) {

	// test 2 account active
	var transRepo TransactionRepository = NewTransactionRepositoryImpl()
	trans := entity.Transaction{AccountNoRsc: 1, AccountNoDes: 2, Amount: 10, Message: "cam on 10k"}
	// fmt.Print(transRepo.Create(trans))
	if err := transRepo.Create(&trans); err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(trans)
	}
}
