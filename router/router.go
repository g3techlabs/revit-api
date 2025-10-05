package router

import (
	"github.com/g3techlabs/revit-api/core/auth/services"
	"github.com/g3techlabs/revit-api/core/mail"
	"github.com/g3techlabs/revit-api/core/token"
	"github.com/g3techlabs/revit-api/core/users/repository"
	"github.com/g3techlabs/revit-api/core/users/service"
	"github.com/g3techlabs/revit-api/validation"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	validator := validation.NewValidator()

	userRepo := repository.NewUserRepository()

	emailService := mail.NewEmailService()
	tokenService := token.NewTokenService()
	authService := services.NewAuthService(validator, userRepo, emailService, tokenService)
	userService := service.NewUserService(validator, userRepo)

	api := app.Group("/api")
	AuthRoutes(api, authService)
	UserRoutes(api, userService, userRepo, tokenService)
}
