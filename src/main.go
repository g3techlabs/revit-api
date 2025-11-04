package main

import (
	"github.com/g3techlabs/revit-api/src/config"
	"github.com/g3techlabs/revit-api/src/db"
	"github.com/g3techlabs/revit-api/src/db/extensions"
	"github.com/g3techlabs/revit-api/src/db/seeders"
	"github.com/g3techlabs/revit-api/src/router"
	"github.com/g3techlabs/revit-api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(config.NewFiberConfig())

	logger := utils.NewLogger()

	SetupDB(logger)

	router.SetupRoutes(app, logger)

	logger.Fatal(app.Listen(":3000"))
}

func SetupDB(logger utils.ILogger) {
	db.Connect(logger)
	extensions.EnableAll(logger)
	db.Migrations(logger)
	seeders.ExecuteAll(logger)
}
