package main

import (
	"github.com/g3techlabs/revit-api/src/config"
	"github.com/g3techlabs/revit-api/src/db"
	"github.com/g3techlabs/revit-api/src/db/extensions"
	"github.com/g3techlabs/revit-api/src/db/seeders"
	"github.com/g3techlabs/revit-api/src/router"
	"github.com/g3techlabs/revit-api/src/utils"
	"github.com/gofiber/fiber/v2"

	_ "github.com/g3techlabs/revit-api/docs" // Import para registrar docs gerados pelo swag
	"github.com/gofiber/swagger"
)

// @title Revit API
// @version 1.0
// @description API do Revit integrada com Swagger via Fiber.
// @BasePath /
// @host localhost:3000
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Digite "Bearer " seguido do token JWT. Exemplo: "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
func main() {
	app := fiber.New(config.NewFiberConfig())

	logger := utils.NewLogger()

	SetupDB(logger)

	// Rota Swagger: http://localhost:3000/swagger/index.html
	app.Get("/swagger/*", swagger.HandlerDefault)

	router.SetupRoutes(app, logger)

	logger.Fatal(app.Listen(":3000"))
}

func SetupDB(logger utils.ILogger) {
	db.Connect(logger)
	extensions.EnableAll(logger)
	db.Migrations(logger)
	seeders.ExecuteAll(logger)
}
