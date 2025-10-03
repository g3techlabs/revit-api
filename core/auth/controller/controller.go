package controller

import (
	"github.com/g3techlabs/revit-api/core/auth/input"
	"github.com/g3techlabs/revit-api/core/auth/middleware"
	"github.com/g3techlabs/revit-api/core/auth/services"
	usersInput "github.com/g3techlabs/revit-api/core/users/input"
	"github.com/g3techlabs/revit-api/utils"
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
	var input usersInput.CreateUser

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body:" + err.Error(),
		})
	}

	if errors := utils.ValidateStruct(input); len(errors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
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
	var input input.LoginCredentials

	if err := ctx.BodyParser(&input); err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body" + err.Error(),
		})
	}

	if errors := utils.ValidateStruct(input); len(errors) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
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
	var input input.Identifier

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	if errs := utils.ValidateStruct(input); len(errs) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errs,
		})
	}

	err := c.AuthService.SendPassResetEmail(input)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c AuthController) ResetPassword(ctx *fiber.Ctx) error {
	var input input.ResetPassword

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	if errs := utils.ValidateStruct(input); len(errs) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errs,
		})
	}

	if err := c.AuthService.ResetPassword(input); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
