package db

import (
	"github.com/g3techlabs/revit-api/db/models"
	"github.com/g3techlabs/revit-api/utils"
)

func Migrations() {
	utils.Log.Info("Initiating database migration...")
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
	)
	if err != nil {
		utils.Log.Fatalf("Error during migrations: %v", err)
	}
	utils.Log.Info("Database successfully migrated.")
}
