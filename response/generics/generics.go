package generics

import "github.com/gofiber/fiber/v2"

func Conflict(message string) error {
	return fiber.NewError(fiber.StatusConflict, message)
}

func InternalError() error {
	return fiber.NewError(fiber.StatusInternalServerError, "Something went wrong unexpectedly")
}

func Unauthorized(message string) error {
	return fiber.NewError(fiber.StatusUnauthorized, message)
}
