package respository

import (
	"gorm.io/gorm"
	models "internBE.com/model"
)

type UserRepository interface {
	CreateBook(user models.User) models.User
	UpdateUsers(user models.User) models.User
	DeleteUserById(userId uint) models.User
	GetAllUsers() []models.User
	FindUserByID(userId uint) models.User
}
type UserConnect struct {
	DB *gorm.DB
}

func PUserConnect(DB *gorm.DB) UserConnect {
	return UserConnect{DB: DB}
}

// create
func (p *UserConnect) CreateUsers(models.User) models.User {
	user := models.User{}

	p.DB.Create(&user)

	return user
}

func (p *UserConnect) GetAllUsers() []models.User {
	var users []models.User
	p.DB.Find(&users)
	return users
}

func (c *UserConnect) FindUserByID(userId uint) models.User {
	var user models.User
	c.DB.First(&user, userId)

	return user
}

func (p *UserConnect) UpdateUsers(user models.User) models.User {
	p.DB.Save(&user).Find(&user)
	// p.DB.Preload("Account").
	return user
}

func (p *UserConnect) DeleteUserById(userId int) error {
	if err := p.DB.Delete(models.User{}, userId).Error; err != nil {
		return err
	}

	return nil
}
