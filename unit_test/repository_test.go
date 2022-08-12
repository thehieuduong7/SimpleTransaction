package unit_test

import (
	"testing"

	models "internBE.com/entity"
	"internBE.com/repository"
	"internBE.com/storage"
)

var (
	repo repository.AccountRepository = repository.NewAccountRepository(storage.GetDB())
)

func TestAccountRepository(t *testing.T) {
	t.Run("test_create_account", func(t *testing.T) {
		repo.CreateAccount(&models.Account{UserId: 1, Surplus: 100})
	})
	t.Run("test_get_account_by_id", func(t *testing.T) {
		repo.GetAccountByUserId(1)

	})
	t.Run("test_delete_account", func(t *testing.T) {
		repo.DeleteAccount(3)
		var account = repo.GetAccountById(3)
		if account.IsActive == true {
			t.Fail()
		}
	})
	t.Run("get by id ", func(t *testing.T) {
		repo.GetAccountById(6)
	})

}
