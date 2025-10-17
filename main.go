package main

import (
	"log"

	"github.com/g3techlabs/revit-api/config"
	"github.com/g3techlabs/revit-api/db"
	"github.com/g3techlabs/revit-api/db/extensions"
	"github.com/g3techlabs/revit-api/db/seeders"
	"github.com/g3techlabs/revit-api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(config.NewFiberConfig())

	SetupDB()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func SetupDB() {
	db.Connect()
	extensions.EnableAll()
	db.Migrations()
	seeders.ExecuteAll()
}
