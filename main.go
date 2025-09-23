package main

import (
	"log"

	"github.com/g3techlabs/revit-api/config"
	"github.com/g3techlabs/revit-api/db"
	"github.com/g3techlabs/revit-api/router"
	"github.com/g3techlabs/revit-api/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadDotEnv()

	utils.Validator.RegisterValidation("uperandlowerrunes", utils.HasUperAndLowerCase)

	app := fiber.New()

	db.Connect()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
