package seeders

import (
	"github.com/g3techlabs/revit-api/db/models"
	"github.com/g3techlabs/revit-api/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func seedRoleTable(db *gorm.DB) {
	utils.Log.Info("Starting Role Table seed...")
	roles := []models.Role{
		{Name: "owner", ID: 1},
		{Name: "admin", ID: 2},
		{Name: "member", ID: 3},
	}

	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&roles).Error; err != nil {
		utils.Log.Errorf("Error seeding Role Table: %v", err)
	}
	utils.Log.Info("Role Table seed complete.")
}
