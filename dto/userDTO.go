package dto

import "errors"

//UserUpdateDTO is used by client when PUT update profile
type UserUpdateDTO struct {
	UserId      int    `json:"user_id"`
	Name        string `json:"name"  binding:"required"`
	PhoneNumber string `json:"phone_number"  binding:"required"`
	Email       string `json:"email"  binding:"required"`
}

type UserCreateDTO struct {
	Name        string `json:"name"  binding:"required"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type UserDTO struct {
	UserId      int    `json:"user_id"`
	Name        string `json:"name"  binding:"required"`
	PhoneNumber string `json:"phone_number"  binding:"required"`
	Email       string `json:"email"  binding:"required"`
}

func (user *UserDTO) CheckExisted() error {

	if user.UserId == 0 {
		return errors.New("not existed this User ")
	}
	return nil
}
