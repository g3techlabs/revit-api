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
	authController := controller.NewUserController(us)

	user := router.Group("/user")

	user.Patch("/", middleware.Auth(userRepository, ts), authController.UpdateUser)
}
