package http

import "github.com/gofiber/fiber/v2"

func NoContent(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusNoContent)
}

func Unauthorized(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": message,
	})
}

func BadRequest(ctx *fiber.Ctx, message string) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": message,
	})
}

func InternalError(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusInternalServerError)
}
