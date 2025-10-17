package seeders

import (
	"github.com/g3techlabs/revit-api/db/models"
	"github.com/g3techlabs/revit-api/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func seedInviteStatusTable(db *gorm.DB) {
	utils.Log.Info("Starting InviteStatus Table seed...")
	statuses := []models.InviteStatus{
		{Status: "accepted", ID: 1},
		{Status: "pending", ID: 2},
		{Status: "rejected", ID: 3},
	}

	if err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&statuses).Error; err != nil {
		utils.Log.Errorf("Error seeding InviteStatus Table: %v", err)
	}
	utils.Log.Info("InviteStatus Table seed complete.")
}
