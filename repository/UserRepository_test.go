package repository

import (
	"testing"

	models "internBE.com/model"
	"internBE.com/storage"
)

var (
	config = &storage.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "123",
		DBName:   "demo1"}
)

func Test(t *testing.T) {

	var test UserRepository = NewUserRepository(storage.Connect(config))

	t.Run("createUsers", func(t *testing.T) {
		user := models.User{Name: "hello", PhoneNumber: "1231441", Email: "hello@example.com"}
		test.CreateUsers(user)
	})

	t.Run("UpdateUsers", func(t *testing.T) {
		user := models.User{Name: "hello"}
		test.UpdateUsers(user)
	})
	t.Run("GetAllUsers", func(t *testing.T) {
		//user := []models.User{}
		test.GetAllUsers()
	})

	t.Run("FindUserByID", func(t *testing.T) {

		test.FindUserByID(1)
	})

	t.Run("DeleteUserById", func(t *testing.T) {

		test.DeleteUserById(1)
	})

	// call router
	// routes.UserRoute(router)
	// router.Run(":8000")

}
