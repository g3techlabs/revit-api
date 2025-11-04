package seeders

import (
	"github.com/g3techlabs/revit-api/src/db/models"
	"github.com/g3techlabs/revit-api/src/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func seedRoleTable(db *gorm.DB, logger utils.ILogger) {
	logger.Info("Starting Role Table seed...")
	roles := []models.Role{
		{Name: "owner", ID: 1},
		{Name: "admin", ID: 2},
		{Name: "member", ID: 3},
	}

	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&roles).Error; err != nil {
		logger.Errorf("Error seeding Role Table: %v", err)
	}
	logger.Info("Role Table seed complete.")
}
