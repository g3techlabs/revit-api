package router

import (
	"github.com/g3techlabs/revit-api/src/core/auth/middleware"
	"github.com/g3techlabs/revit-api/src/core/users/controller"
	"github.com/g3techlabs/revit-api/src/core/users/service"
	"github.com/g3techlabs/revit-api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router, us service.IUserService, middleware *middleware.AuthMiddleware, logger utils.ILogger) {
	logger.Info("Setting up USER routes...")

	userController := controller.NewUserController(us)

	user := router.Group("/user", middleware.Auth())
	friendship := user.Group("/friendship")

	friendship.Get("/", userController.GetFriends)
	friendship.Get("/requests", userController.GetFriendshipRequests)
	friendship.Post("/:destinataryId", userController.RequestFriendship)
	friendship.Patch("/:requesterId", userController.AnswerFriendshipRequest)
	friendship.Delete("/:friendId", userController.RemoveFriendship)

	user.Get("/", userController.GetUsers)
	user.Get("/email-available", userController.CheckIfEmailAvailable)
	user.Get("/nickname-available", userController.CheckIfNicknameAvailable)
	user.Get("/:id", userController.GetUser)
	user.Patch("/", userController.UpdateUser)
	user.Post("/profile-pic/", userController.RequestProfilePicUpdate)
	user.Patch("/profile-pic", userController.ConfirmNewProfilePic)

	logger.Info("USER routes successfully set up.")
}
