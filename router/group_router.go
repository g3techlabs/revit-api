package router

import (
	"github.com/g3techlabs/revit-api/core/auth/middleware"
	"github.com/g3techlabs/revit-api/core/group/controller"
	"github.com/g3techlabs/revit-api/core/group/service"
	"github.com/g3techlabs/revit-api/utils"
	"github.com/gofiber/fiber/v2"
)

func GroupRoutes(router fiber.Router, groupService service.IGroupService, middleware *middleware.AuthMiddleware) {
	utils.Log.Info("Setting up GROUP routes...")

	groupController := controller.NewGroupController(groupService)

	group := router.Group("/group", middleware.Auth())
	group.Post("/", groupController.CreateGroup)
	group.Get("/", groupController.GetGroups)
	group.Get("/invite", groupController.GetPendingInvites)

	photos := router.Group("/photos")
	photos.Put("/:groupId", groupController.RequestNewGroupPhotos)
	photos.Patch("/:groupId", groupController.ConfirmNewPhotos)

	group.Post("/:groupId/member", groupController.JoinGroup)
	group.Delete("/:groupId/member", groupController.QuitGroup)
	group.Post("/:groupId/invite/:invitedId", groupController.InviteUser)
	group.Patch("/:groupId/invite", groupController.AnswerPendingInvite)

	group.Patch("/:groupId", groupController.UpdateGroup)

	utils.Log.Info("GROUP routes successfully set up.")
}
