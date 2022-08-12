package dto

//UserUpdateDTO is used by client when PUT update profile
type UserUpdateDTO struct {
	UserId      int    `json:"user_id"`
	Name        string `json:"name"  binding:"required"`
	PhoneNumber string `json:"phone_number"  binding:"required"`
	Email       string `json:"email"  binding:"required"`
}

type UserCreateDTO struct {
	Name        string `json:"name"  binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}
