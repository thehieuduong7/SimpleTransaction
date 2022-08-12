package unit_test

import (
	"testing"

	models "internBE.com/entity"
	"internBE.com/repository"
	"internBE.com/storage"
)

var (
	test repository.UserRepository = repository.NewUserRepository(storage.Connect())
)

func Test(t *testing.T) {

	//  var test UserRepository = NewUserRepository(storage.Connect())

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

		test.DeleteUser(1)
	})

	// call router
	// routes.UserRoute(router)
	// router.Run(":8000")

}
