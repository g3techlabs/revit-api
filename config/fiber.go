package config

import (
	"errors"

	"github.com/g3techlabs/revit-api/response"
	"github.com/g3techlabs/revit-api/validation"
	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		AppName:      Get("APP_NAME") + " API",
		ErrorHandler: ErrorHandler,
	}
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	if errorsMap := validation.ValidationErrorMessages(err); len(errorsMap) > 0 {
		return response.Error(c, fiber.StatusBadRequest, "Bad Request", errorsMap)
	}

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		return response.Error(c, fiberErr.Code, fiberErr.Message, nil)
	}

	var customErr *response.CustomError
	if errors.As(err, &customErr) {
		return response.Error(c, customErr.StatusCode, customErr.Message, customErr.Details)
	}

	return response.Error(c, fiber.StatusInternalServerError, "Internal Server Error", nil)
}
