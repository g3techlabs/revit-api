package handler

import (
	"github.com/g3techlabs/revit-api/core/auth/services"
	"github.com/gofiber/fiber/v2"
)

func RegisterUser(ctx *fiber.Ctx) error {
	return services.RegisterUser(ctx)
}

func Login(ctx *fiber.Ctx) error {
	return services.Login(ctx)
}

func RefreshTokens(ctx *fiber.Ctx) error {
	return services.RefreshTokens(ctx)
}
