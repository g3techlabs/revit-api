package router

import (
	"github.com/g3techlabs/revit-api/core/auth/middleware"
	"github.com/g3techlabs/revit-api/core/token"
	"github.com/g3techlabs/revit-api/core/users/controller"
	"github.com/g3techlabs/revit-api/core/users/repository"
	"github.com/g3techlabs/revit-api/core/users/service"
	"github.com/g3techlabs/revit-api/utils"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router, us service.IUserService, userRepository repository.UserRepository, ts token.ITokenService) {
	utils.Log.Info("Setting up USER routes...")

	userController := controller.NewUserController(us)

	user := router.Group("/user", middleware.Auth(userRepository, ts))
	friendship := user.Group("/friendship")

	friendship.Get("/", userController.GetFriends)
	friendship.Post("/:destinataryId", userController.RequestFriendship)
	friendship.Patch("/:requesterId", userController.AnswerFriendshipRequest)
	friendship.Delete("/:friendId", userController.RemoveFriendship)

	user.Get("/", userController.GetUsers)
	user.Get("/:id", userController.GetUser)
	user.Patch("/", userController.UpdateUser)
	user.Post("/profile-pic/", userController.RequestProfilePicUpdate)
	user.Patch("/profile-pic", userController.ConfirmNewProfilePic)

	utils.Log.Info("USER routes successfully set up.")
}
