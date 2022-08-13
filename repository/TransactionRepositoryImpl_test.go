package repository

import (
	"testing"

	"internBE.com/entity"
	"internBE.com/storage"
)

func Test_transactionRepositoryImpl_Create(t *testing.T) {
	type args struct {
		trans *entity.Transaction
	}
	tests := []struct {
		name    string
		tr      TransactionRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "1",
			tr:      NewTransactionRepositoryImpl(storage.GetDB()),
			args:    args{trans: &entity.Transaction{AccountNoRsc: 2, AccountNoDes: 3, Amount: 3, Message: "hello"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.tr.Create(tt.args.trans); (err != nil) != tt.wantErr {
				t.Errorf("transactionRepositoryImpl.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
