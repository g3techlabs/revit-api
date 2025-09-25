package router

import (
	"github.com/g3techlabs/revit-api/core/auth/handler"
	"github.com/g3techlabs/revit-api/core/auth/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", handler.RegisterUser)
	auth.Post("/login", handler.Login)
	auth.Get("/test", middleware.JWTAuth(), func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "properly working 👌",
		})
	})
}
