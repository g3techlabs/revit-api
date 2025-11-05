package db

import (
	"github.com/g3techlabs/revit-api/src/db/models"
	"github.com/g3techlabs/revit-api/src/utils"
)

func Migrations(logger utils.ILogger) {
	logger.Info("Initiating database migration...")
	err := Db.AutoMigrate(
		&models.InviteStatus{},
		&models.Friendship{},
		&models.User{},
		&models.Vehicle{},
		&models.Photo{},
		&models.Country{},
		&models.State{},
		&models.City{},
		&models.Role{},
		&models.Visibility{},
		&models.Group{},
		&models.GroupMember{},
		&models.Event{},
		&models.EventSubscriber{},
		&models.Route{},
		&models.RouteParticipant{},
	)
	if err != nil {
		logger.Fatalf("Error during migrations: %v", err)
	}
	logger.Info("Database successfully migrated.")
}
