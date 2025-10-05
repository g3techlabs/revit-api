package main

import (
	"log"

	"github.com/g3techlabs/revit-api/config"
	"github.com/g3techlabs/revit-api/db"
	"github.com/g3techlabs/revit-api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(config.NewFiberConfig())

	db.Connect()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
