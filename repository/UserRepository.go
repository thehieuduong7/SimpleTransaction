package repository

import (
	// "gorm.io/gorm"
	models "internBE.com/entity"
)

type UserRepository interface {
	CreateUsers(user models.User) models.User
	UpdateUsers(user models.User) models.User

	GetAllUsers() []models.User
	FindUserByID(userId int) models.User
	Save(user models.User) models.User
	DeleteUser(userId int) models.User
}
