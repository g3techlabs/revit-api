package controller

import (
	"strconv"

	"github.com/g3techlabs/revit-api/src/core/users/input"
	"github.com/g3techlabs/revit-api/src/core/users/service"
	"github.com/g3techlabs/revit-api/src/response/generics"
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

func (uc *UserController) RequestProfilePicUpdate(ctx *fiber.Ctx) error {
	input := new(input.RequestProfilePicUpdate)

	if err := ctx.BodyParser(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := uc.userService.RequestProfilePicUpdate(userId, input)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (uc *UserController) ConfirmNewProfilePic(ctx *fiber.Ctx) error {
	input := new(input.ConfirmNewProfilePic)

	if err := ctx.BodyParser(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := uc.userService.ConfirmNewProfilePic(userId, input); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (uc *UserController) GetUsers(ctx *fiber.Ctx) error {
	query := new(input.GetUsersQuery)

	if err := ctx.QueryParser(query); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid query parameters: " + err.Error(),
		})
	}

	users, err := uc.userService.GetUsers(query)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(users)
}

func (uc *UserController) GetUser(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")

	idUint64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}
	id := uint(idUint64)

	response, err := uc.userService.GetUser(id)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (uc *UserController) GetFriends(ctx *fiber.Ctx) error {
	query := new(input.GetUsersQuery)

	if err := ctx.QueryParser(query); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid query parameters: " + err.Error(),
		})
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	friends, err := uc.userService.GetFriends(userId, query)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(friends)
}

func (uc *UserController) RequestFriendship(ctx *fiber.Ctx) error {
	destinataryParam := ctx.Params("destinataryId")

	destinataryUint64, err := strconv.ParseUint(destinataryParam, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid destinatary ID")
	}
	destinataryId := uint(destinataryUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := uc.userService.RequestFriendship(userId, destinataryId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (uc *UserController) AnswerFriendshipRequest(ctx *fiber.Ctx) error {
	requesterParam := ctx.Params("requesterId")

	requesterUint64, err := strconv.ParseUint(requesterParam, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid requester ID")
	}
	requesterId := uint(requesterUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	answer := new(input.FriendshipRequestAnswer)
	if err := ctx.BodyParser(answer); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	response, err := uc.userService.AnswerFriendshipRequest(userId, requesterId, answer)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusNoContent).JSON(response)
}

func (uc *UserController) RemoveFriendship(ctx *fiber.Ctx) error {
	friendParam := ctx.Params("friendId")

	friendUint64, err := strconv.ParseUint(friendParam, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid friend ID")
	}
	friendId := uint(friendUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := uc.userService.RemoveFriendship(userId, friendId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (uc *UserController) GetFriendshipRequests(ctx *fiber.Ctx) error {
	query := new(input.GetFriendshipRequestsQuery)

	if err := ctx.QueryParser(query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := uc.userService.GetFriendshipRequests(userId, query)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *UserController) CheckIfEmailAvailable(ctx *fiber.Ctx) error {
	var input input.EmailInput

	if err := ctx.QueryParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	isEmailAvailable, err := c.userService.CheckIfEmailAvailable(&input)
	if err != nil {
		return err
	}

	if !isEmailAvailable {
		return fiber.NewError(fiber.StatusConflict, "Email already taken")
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *UserController) CheckIfNicknameAvailable(ctx *fiber.Ctx) error {
	var input input.NicknameInput

	if err := ctx.QueryParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	isEmailAvailable, err := c.userService.CheckIfNicknameAvailable(&input)
	if err != nil {
		return err
	}

	if !isEmailAvailable {
		return fiber.NewError(fiber.StatusConflict, "Nickname already taken")
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
