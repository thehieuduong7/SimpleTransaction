package controller

import (
	models "internBE.com/model"
	"internBE.com/repository"
)

type AccountController interface {
	CreatAccount(account *models.Account) *models.Account
	UpdateAccount(account *models.Account) models.Account
	DeleteAccount(id int)
	GetAccountByUserId(id int) []models.Account
	GetAccountById(id int) models.Account
}
type accountController struct {
	AccountRepo repository.AccountRepository
}

func NewAccountController(repo repository.AccountRepository) AccountController {
	return &accountController{AccountRepo: repo}
}

func (controller *accountController) CreatAccount(account *models.Account) *models.Account {
	return controller.AccountRepo.CreatAccount(account)
}

func (controller *accountController) UpdateAccount(account *models.Account) models.Account {
	return controller.AccountRepo.UpdateAccount(account)
}

func (controller *accountController) DeleteAccount(id int) {
	controller.AccountRepo.DeleteAccount(id)
}

func (controller *accountController) GetAccountByUserId(id int) []models.Account {
	return controller.AccountRepo.GetAccountByUserId(id)
}

func (controller *accountController) GetAccountById(id int) models.Account {
	return controller.AccountRepo.GetAccountById(id)
}
