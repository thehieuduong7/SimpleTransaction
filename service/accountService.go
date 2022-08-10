package service

import (
	models "internBE.com/model"
	"internBE.com/repository"
)

type AccountService interface {
	CreateAccount(account *models.Account) *models.Account
	UpdateAccount(account *models.Account) models.Account
	DeleteAccount(id int)
	GetAccountByUserId(id int) []models.Account
	GetAccountById(id int) models.Account
}
type accountService struct {
	AccountRepo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) AccountService {
	return &accountService{AccountRepo: repo}
}

func (controller *accountService) CreateAccount(account *models.Account) *models.Account {
	return controller.AccountRepo.CreateAccount(account)
}

func (controller *accountService) UpdateAccount(account *models.Account) models.Account {
	return controller.AccountRepo.UpdateAccount(account)
}

func (controller *accountService) DeleteAccount(id int) {
	controller.AccountRepo.DeleteAccount(id)
}

func (controller *accountService) GetAccountByUserId(id int) []models.Account {
	return controller.AccountRepo.GetAccountByUserId(id)
}

func (controller *accountService) GetAccountById(id int) models.Account {
	return controller.AccountRepo.GetAccountById(id)
}
