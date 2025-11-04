package seeders

import (
	"github.com/g3techlabs/revit-api/src/db/models"
	"github.com/g3techlabs/revit-api/src/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func seedVisibilityTable(db *gorm.DB, logger utils.ILogger) {
	logger.Info("Starting Visibility Table seed...")
	visibilities := []models.Visibility{
		{Name: "public", ID: 1},
		{Name: "private", ID: 2},
	}

	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&visibilities).Error; err != nil {
		logger.Errorf("Error seeding Visibility Table: %v", err)
	}
	logger.Info("Visibility Table seed complete.")
}
