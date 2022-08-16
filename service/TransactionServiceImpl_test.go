package service

import (
	"errors"
	"fmt"
	"testing"

	"internBE.com/repository"
	"internBE.com/storage"
)

func Test_transactionServiceImpl_GetStaticTransaction(t *testing.T) {
	// req := dto.StaticTransactionRequest{AccountNo1: 0, AccountNo2: 3}
	// repo := repository.
	// 	NewTransactionRepositoryImpl(storage.GetDB())
	// service := NewTransactionServiceImpl(&repo)
	// res, err := service.GetStaticTransaction(req)
	// if err != nil {
	// 	t.Log(err.Error())
	// 	t.Fail()
	// } else {
	// 	t.Log(res)
	// }

	accRepo := repository.NewAccountRepository(storage.GetDB())
	if !accRepo.IsExistAccount(7) {
		msg := fmt.Sprintf("account no %d not found", 0)
		t.Log(errors.New(msg).Error())
		t.Fail()
		return
	}
	t.Log("succeess")
}
