package router

import (
	"github.com/g3techlabs/revit-api/src/core/auth/controller"
	"github.com/g3techlabs/revit-api/src/core/auth/services"
	"github.com/g3techlabs/revit-api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router, as services.IAuthService, logger utils.ILogger) {
	logger.Info("Setting up AUTH routes...")

	authController := controller.NewAuthController(as)

	auth := router.Group("/auth")

	auth.Get("/email-available", authController.CheckIfEmailAvailable)
	auth.Get("/nickname-available", authController.CheckIfNicknameAvailable)
	auth.Post("/register", authController.RegisterUser)
	auth.Post("/login", authController.Login)
	auth.Post("/refresh", authController.RefreshTokens)
	auth.Post("/password", authController.SendPassResetEmail)
	auth.Patch("/password", authController.ResetPassword)

	logger.Info("AUTH routes successfully set up.")
}
