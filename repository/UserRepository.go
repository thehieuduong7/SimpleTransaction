package repository

import (
	"gorm.io/gorm"
	models "internBE.com/model"
)

type UserRepository interface {
	CreateUsers(user models.User) models.User
	UpdateUsers(user models.User) models.User
	DeleteUserById(userId uint) models.User
	GetAllUsers() []models.User
	FindUserByID(userId uint) models.User
}
type userConnection struct {
	DB *gorm.DB
}

func NewUserRepository(dbConn *gorm.DB) UserRepository {
	return &userConnection{DB: dbConn}
}

// create
func (p *userConnection) CreateUsers(user models.User) models.User {
	p.DB.Save(&user)
	p.DB.Create(&user)

	return user
}

func (p *userConnection) GetAllUsers() []models.User {
	var users []models.User
	p.DB.Find(&users)
	return users
}

func (c *userConnection) FindUserByID(userId uint) models.User {
	var user models.User
	c.DB.Where("id = ?", user.UserId).Take(&user)
	return user
}

func (p *userConnection) UpdateUsers(user models.User) models.User {
	p.DB.Save(&user).Where("name=?", user.Name).Updates(user)
	// p.DB.Preload("Account").
	return user
}

func (p *userConnection) DeleteUserById(userId uint) models.User {
	var user models.User
	p.DB.Delete(&user, userId)
	return user

}
