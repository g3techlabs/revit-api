package controller

import (
	"strconv"

	"github.com/g3techlabs/revit-api/src/core/group/input"
	_ "github.com/g3techlabs/revit-api/src/core/group/response"
	"github.com/g3techlabs/revit-api/src/core/group/service"
	"github.com/g3techlabs/revit-api/src/response/generics"
	"github.com/gofiber/fiber/v2"
)

type GroupController struct {
	groupService service.IGroupService
}

func NewGroupController(groupService service.IGroupService) *GroupController {
	return &GroupController{groupService: groupService}
}

func (c *GroupController) CreateGroup(ctx *fiber.Ctx) error {
	data := new(input.CreateGroup)

	if err := ctx.BodyParser(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := c.groupService.CreateGroup(userId, data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *GroupController) ConfirmNewPhotos(ctx *fiber.Ctx) error {
	data := new(input.ConfirmNewPhotos)

	if err := ctx.BodyParser(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	groupParam := ctx.Params("groupId")

	groupIdUint64, err := strconv.ParseUint(groupParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}
	groupId := uint(groupIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.groupService.ConfirmNewPhotos(userId, groupId, data); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *GroupController) GetGroups(ctx *fiber.Ctx) error {
	query := new(input.GetGroupsQuery)

	if err := ctx.QueryParser(query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := c.groupService.GetGroups(userId, query)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *GroupController) GetGroup(ctx *fiber.Ctx) error {
	groupParam := ctx.Params("groupId")

	groupIdUint64, err := strconv.ParseUint(groupParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}
	groupId := uint(groupIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := c.groupService.GetGroup(userId, groupId)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *GroupController) GetMembers(ctx *fiber.Ctx) error {
	query := new(input.GetMembersInput)

	if err := ctx.QueryParser(query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	groupParam := ctx.Params("groupId")

	groupIdUint64, err := strconv.ParseUint(groupParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}
	groupId := uint(groupIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := c.groupService.GetMembers(userId, groupId, query)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *GroupController) UpdateGroup(ctx *fiber.Ctx) error {
	data := new(input.UpdateGroup)

	if err := ctx.BodyParser(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	groupParam := ctx.Params("groupId")
	groupIdUint64, err := strconv.ParseUint(groupParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}
	groupId := uint(groupIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.groupService.UpdateGroup(userId, groupId, data); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *GroupController) RequestNewGroupPhotos(ctx *fiber.Ctx) error {
	data := new(input.RequestNewGroupPhotos)

	if err := ctx.BodyParser(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	groupParam := ctx.Params("groupId")
	groupIdUint64, err := strconv.ParseUint(groupParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}
	groupId := uint(groupIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := c.groupService.RequestNewGroupPhotos(userId, groupId, data)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *GroupController) JoinGroup(ctx *fiber.Ctx) error {
	groupParam := ctx.Params("groupId")
	groupIdUint64, err := strconv.ParseUint(groupParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}
	groupId := uint(groupIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.groupService.JoinGroup(userId, groupId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *GroupController) QuitGroup(ctx *fiber.Ctx) error {
	groupParam := ctx.Params("groupId")
	groupIdUint64, err := strconv.ParseUint(groupParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}
	groupId := uint(groupIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.groupService.QuitGroup(userId, groupId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *GroupController) InviteUser(ctx *fiber.Ctx) error {
	groupParam := ctx.Params("groupId")
	groupIdUint64, err := strconv.ParseUint(groupParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}
	groupId := uint(groupIdUint64)

	invitedParam := ctx.Params("invitedId")
	invitedIdUint64, err := strconv.ParseUint(invitedParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid invited user ID",
		})
	}
	invitedId := uint(invitedIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.groupService.InviteUser(userId, groupId, invitedId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *GroupController) GetPendingInvites(ctx *fiber.Ctx) error {
	var query input.GetPendingInvites

	if err := ctx.QueryParser(&query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query parameters")
	}

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	response, err := c.groupService.GetPendingInvites(userId, &query)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *GroupController) AnswerPendingInvite(ctx *fiber.Ctx) error {
	var answer input.AnswerPendingInvite

	if err := ctx.BodyParser(&answer); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid body request")
	}

	groupParam := ctx.Params("groupId")
	groupIdUint64, err := strconv.ParseUint(groupParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}
	groupId := uint(groupIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.groupService.AnswerPendingInvite(userId, groupId, &answer); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *GroupController) RemoveMember(ctx *fiber.Ctx) error {
	groupParam := ctx.Params("groupId")
	groupIdUint64, err := strconv.ParseUint(groupParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}
	groupId := uint(groupIdUint64)

	memberParam := ctx.Params("memberId")
	memberIdUint64, err := strconv.ParseUint(memberParam, 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid member ID",
		})
	}
	memberId := uint(memberIdUint64)

	userId, ok := ctx.Locals("userId").(uint)
	if !ok {
		return generics.Unauthorized("Invalid or non-existent auth token")
	}

	if err := c.groupService.RemoveMember(userId, groupId, memberId); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
