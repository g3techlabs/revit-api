package seeders

import (
	"github.com/g3techlabs/revit-api/db/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func seedInviteStatusTable(db *gorm.DB) error {
	statuses := []models.InviteStatus{
		{Status: "accepted", ID: 1},
		{Status: "pending", ID: 2},
		{Status: "rejected", ID: 3},
	}

	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&statuses).Error
}
