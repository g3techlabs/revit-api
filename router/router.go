package router

import (
	"github.com/g3techlabs/revit-api/core/auth/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/auth")
	api.Post("/register", handler.RegisterUser)
	api.Get("/login", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Logou")
	})
}
