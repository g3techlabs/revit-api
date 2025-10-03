package router

import (
	"github.com/g3techlabs/revit-api/core/auth/services"
	"github.com/g3techlabs/revit-api/core/mail"
	"github.com/g3techlabs/revit-api/core/token"
	"github.com/g3techlabs/revit-api/core/users/repository"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	userRepo := repository.NewUserRepository()

	emailService := mail.NewEmailService()
	tokenService := token.NewTokenService()
	authService := services.NewAuthService(userRepo, emailService, tokenService)

	api := app.Group("/api")
	AuthRoutes(api, authService)
}
