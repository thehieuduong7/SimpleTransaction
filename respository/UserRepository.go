package respository

import (
	"gorm.io/gorm"
	models "internBE.com/model"
)

type UserRepository interface {
	// CreateBook(b models.User) models.User
	UpdateUsers(b models.User) models.User
	DeleteUserById(b models.User)
	GetAllUsers() []models.User
	FindUserByID(UserId uint) *models.User
}
type UserConnect struct {
	DB *gorm.DB
}

func PUserConnect(DB *gorm.DB) UserConnect {
	return UserConnect{DB: DB}
}

// create
// func (p *UserConnect) CreateUsers(b models.User) (models.User, error) {
// 	user := &models.User{}

// 	if err := p.DB.Create(&user).Error; err != nil {
// 		return nil, err
// 	}

// 	return b, nil
// }
func (p *UserConnect) GetAllUsers() []models.User {
	var users []models.User
	p.DB.Preload("Account").Find(&users)
	return users
}

func (c UserConnect) FindUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := c.DB.Preload("Account").First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *UserConnect) UpdateUsers(user models.User) models.User {
	p.DB.Save(&user)
	p.DB.Preload("Account").Find(&user)

	return user
}

func (p *UserConnect) DeleteUserById(id int) error {
	if err := p.DB.Delete(models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
