package router

import (
	"github.com/g3techlabs/revit-api/core/auth/middleware"
	"github.com/g3techlabs/revit-api/core/group/controller"
	"github.com/g3techlabs/revit-api/core/group/service"
	"github.com/g3techlabs/revit-api/utils"
	"github.com/gofiber/fiber/v2"
)

func GroupRoutes(router fiber.Router, groupService service.IGroupService, middleware *middleware.AuthMiddleware) {
	utils.Log.Info("Setting up VEHICLE routes...")

	groupController := controller.NewGroupController(groupService)

	group := router.Group("/group", middleware.Auth())

	group.Post("/", groupController.CreateGroup)
	group.Get("/", groupController.GetGroups)
	group.Patch("/photos/:groupId", groupController.ConfirmNewPhotos)

	utils.Log.Info("VEHICLE routes successfully set up.")
}
