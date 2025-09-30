package router

import (
	"github.com/g3techlabs/revit-api/core/auth/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", handler.RegisterUser)
	auth.Post("/login", handler.Login)
	auth.Post("/refresh", handler.RefreshTokens)
	auth.Post("/send-reset-password-email", handler.SendResetPasswordEmail)
	auth.Post("/reset-password", handler.ResetPassword)
}
