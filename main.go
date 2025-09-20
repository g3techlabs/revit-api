package main

import (
	"github.com/g3techlabs/revit-api/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db.Setup()

	app := fiber.New()

	app.Get("/", func(context *fiber.Ctx) error {
		return context.SendString("Hello world!")
	})

	app.Listen(":3000")
}
