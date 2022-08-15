package repository

import (
	"gorm.io/gorm"
	models "internBE.com/entity"
)

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(dbConn *gorm.DB) UserRepository {
	return &userConnection{connection: dbConn}
}

func (db *userConnection) CreateUsers(user models.User) models.User {
	db.connection.Preload("Accounts").Save(&user).Find(&user)

	return user
}

func (db *userConnection) GetAllUsers() []models.User {
	var users []models.User
	db.connection.Preload("Accounts").Find(&users)
	return users
}

func (db *userConnection) FindUserByID(userId int) models.User {
	var users models.User

	db.connection.Preload("Accounts").Find(&users, userId)

	return users
}

func (db *userConnection) UpdateUsers(user models.User) models.User {

	db.connection.Save(&user).Where(user.UserId).Updates(user)
	return user
}

func (db *userConnection) DeleteUser(userId int) models.User {
	var user models.User
	db.connection.Preload("Accounts").Delete(user, userId)
	return user
}

func (db *userConnection) Save(user models.User) models.User {
	db.connection.Save(&user)

	return user
}
