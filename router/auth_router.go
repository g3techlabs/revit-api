package router

import (
	"github.com/g3techlabs/revit-api/core/auth/controller"
	"github.com/g3techlabs/revit-api/core/auth/services"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router, as services.IAuthService) {
	authController := controller.NewAuthController(as)

	auth := router.Group("/auth")

	auth.Post("/register", authController.RegisterUser)
	auth.Post("/login", authController.Login)
	auth.Post("/refresh", authController.RefreshTokens)
	auth.Post("/send-reset-password-email", authController.SendPassResetEmail)
	auth.Post("/reset-password", authController.ResetPassword)
}
