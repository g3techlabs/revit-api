package router

import (
	"github.com/g3techlabs/revit-api/core/auth/middleware"
	"github.com/g3techlabs/revit-api/core/token"
	"github.com/g3techlabs/revit-api/core/users/controller"
	"github.com/g3techlabs/revit-api/core/users/repository"
	"github.com/g3techlabs/revit-api/core/users/service"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router, us service.IUserService, userRepository repository.UserRepository, ts token.ITokenService) {
	userController := controller.NewUserController(us)

	user := router.Group("/user", middleware.Auth(userRepository, ts))

	user.Patch("/", userController.UpdateUser)
	user.Patch("/profile-pic", userController.UpdateProfilePic)
	user.Post("/profile-pic/presign", userController.PresignProfilePic)
}
