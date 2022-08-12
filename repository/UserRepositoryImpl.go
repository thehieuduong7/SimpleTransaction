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

	user = models.User{}

	db.connection.Save(&user).Create(&user)

	return user
}

func (db *userConnection) GetAllUsers() []models.User {
	var users []models.User
	db.connection.Find(&users)
	return users
}

func (db *userConnection) FindUserByID(userId int) models.User {
	var users models.User

	db.connection.Find(&users, "user_id=? AND is_active= ?", userId)

	return users
}

func (db *userConnection) UpdateUsers(user models.User) models.User {

	db.connection.Save(&user).Where("user_id=? AND is_active= ?", user.UserId).Updates(user)
	return user
}

func (db *userConnection) DeleteUser(userId int) models.User {
	var user models.User
	db.connection.Delete(user, userId)
	return user
}

func (db *userConnection) Save(user models.User) models.User {
	db.connection.Save(&user)

	return user
}
