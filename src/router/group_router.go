package router

import (
	"github.com/g3techlabs/revit-api/src/core/auth/middleware"
	"github.com/g3techlabs/revit-api/src/core/group/controller"
	"github.com/g3techlabs/revit-api/src/core/group/service"
	"github.com/g3techlabs/revit-api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func GroupRoutes(router fiber.Router, groupService service.IGroupService, middleware *middleware.AuthMiddleware, logger utils.ILogger) {
	logger.Info("Setting up GROUP routes...")

	groupController := controller.NewGroupController(groupService)

	group := router.Group("/group", middleware.Auth())
	group.Post("/", groupController.CreateGroup)
	group.Get("/", groupController.GetGroups)
	group.Get("/invite", groupController.GetPendingInvites)
	group.Get("/:groupId", groupController.GetGroup)
	group.Get("/:groupId/member", groupController.GetMembers)

	group.Put("/photos/:groupId", groupController.RequestNewGroupPhotos)
	group.Patch("/photos/:groupId", groupController.ConfirmNewPhotos)

	group.Post("/:groupId/member", groupController.JoinGroup)
	group.Delete("/:groupId/member", groupController.QuitGroup)
	group.Delete("/:groupId/member/:memberId", groupController.RemoveMember)
	group.Post("/:groupId/invite/:invitedId", groupController.InviteUser)
	group.Patch("/:groupId/invite", groupController.AnswerPendingInvite)

	group.Patch("/:groupId", groupController.UpdateGroup)

	logger.Info("GROUP routes successfully set up.")
}
