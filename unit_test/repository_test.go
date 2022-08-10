package unit_test

import (
	"internBE.com/constant"
	models "internBE.com/model"
	"internBE.com/respository"
	"internBE.com/storage"
	"testing"
)

var (
	config = &storage.Config{
		Host:     constant.Host,
		Port:     constant.Port,
		Password: constant.Password,
		User:     constant.User,
		DBName:   constant.DBName}
	repo respository.AccountRepository = respository.NewAccountRepository(storage.Connect(config))
)

func TestAccountRepository(t *testing.T) {
	t.Run("test_create_account", func(t *testing.T) {
		repo.CreatAccount(&models.Account{UserId: 1, Surplus: 100})
	})
	t.Run("test_get_account_by_id", func(t *testing.T) {
		repo.GetAccountByUserId(1)

	})
	t.Run("test_delte_account", func(t *testing.T) {
		repo.DeleteAccount(3)
		var account = repo.GetAccountById(3)
		if account.IsActive == true {
			t.Fail()
		}
	})

}
