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

func Forbidden(message string) error {
	return fiber.NewError(fiber.StatusForbidden, message)
}

func InvalidFileExtension() error {
	return fiber.NewError(fiber.StatusBadRequest, "Invalid file extension. Only .jpg, .jpeg and .png are allowed.")
}
