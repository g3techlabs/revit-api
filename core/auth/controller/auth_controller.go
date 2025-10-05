package controller

import (
	"github.com/g3techlabs/revit-api/core/auth/input"
	"github.com/g3techlabs/revit-api/core/auth/middleware"
	"github.com/g3techlabs/revit-api/core/auth/services"
	usersInput "github.com/g3techlabs/revit-api/core/users/input"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	AuthService services.IAuthService
}

func NewAuthController(authService services.IAuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

func (c *AuthController) RegisterUser(ctx *fiber.Ctx) error {
	input := new(usersInput.CreateUser)

	if err := ctx.BodyParser(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body:" + err.Error(),
		})
	}

	user, err := c.AuthService.RegisterUser(input)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user": user,
	})
}

func (c AuthController) Login(ctx *fiber.Ctx) error {
	input := new(input.LoginCredentials)

	if err := ctx.BodyParser(input); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body" + err.Error(),
		})
	}

	authTokens, err := c.AuthService.Login(input)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(authTokens)
}

func (c AuthController) RefreshTokens(ctx *fiber.Ctx) error {
	refreshToken, err := middleware.ExtractBearerToken(ctx)
	if err != nil {
		return err
	}

	authTokens, err := c.AuthService.RefreshTokens(refreshToken)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(authTokens)
}

func (c AuthController) SendPassResetEmail(ctx *fiber.Ctx) error {
	input := new(input.Identifier)

	if err := ctx.BodyParser(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	err := c.AuthService.SendPassResetEmail(input)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c AuthController) ResetPassword(ctx *fiber.Ctx) error {
	input := new(input.ResetPassword)

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	if err := c.AuthService.ResetPassword(input); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
