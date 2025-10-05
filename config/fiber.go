package config

import (
	"github.com/g3techlabs/revit-api/utils"
	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		AppName:      Get("APP_NAME") + " API",
		ErrorHandler: utils.ErrorHandler,
	}
}
