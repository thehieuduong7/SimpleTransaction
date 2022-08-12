package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"internBE.com/controller"
	"internBE.com/repository"
	"internBE.com/service"
	"internBE.com/storage"
)

var (
	db             *gorm.DB                  = storage.Connect()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	userService    service.UserService       = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService)
)

func UserRoute(router *gin.Engine) {

	userRoutes := router.Group("/api/users")
	{
		// test postman successful
		userRoutes.GET("/", userController.GetAllUsers)
		userRoutes.POST("/create", userController.CreateUsers)
		userRoutes.GET("/:user_id", userController.GetUserByID)
		userRoutes.PUT("/", userController.UpdateUsers)
		userRoutes.DELETE("/:user_id", userController.DeleteUserById)

	}

}
