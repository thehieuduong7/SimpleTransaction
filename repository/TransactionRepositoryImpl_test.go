package repository

import (
	"testing"
	"time"

	"internBE.com/constantGlobal"
	"internBE.com/storage"
)

func Test_transactionRepositoryImpl_Create(t *testing.T) {
	tr := NewTransactionRepositoryImpl(storage.GetDB())
	time1, _ := time.Parse(constantGlobal.TimeLayout, "2022-08-01 20:41:14")
	time2, _ := time.Parse(constantGlobal.TimeLayout, "2022-08-15 20:41:14")
	t.Log(tr.GetHistoryTransBetweenDateWith(2, 1, time1, time2, -1))
}
