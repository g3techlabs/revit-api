package controller

import (
	"github.com/g3techlabs/revit-api/core/users/input"
	"github.com/g3techlabs/revit-api/core/users/service"
	"github.com/g3techlabs/revit-api/response/generics"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) UpdateUser(ctx *fiber.Ctx) error {
	input := new(input.UpdateUser)

	if err := ctx.BodyParser(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := uc.userService.Update(userId, input); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (uc *UserController) PresignProfilePic(ctx *fiber.Ctx) error {
	input := new(input.PresignProfilePic)

	if err := ctx.BodyParser(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := uc.userService.PresignProfilePic(userId, input)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (uc *UserController) UpdateProfilePic(ctx *fiber.Ctx) error {
	input := new(input.UpdateProfilePic)

	if err := ctx.BodyParser(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := uc.userService.UpdateProfilePic(userId, input); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
