package seeders

import (
	"github.com/g3techlabs/revit-api/db"
	"github.com/g3techlabs/revit-api/utils"
)

func ExecuteAll() error {
	utils.Log.Info("Starting InviteStatus Table seed...")
	if err := seedInviteStatusTable(db.Db); err != nil {
		utils.Log.Errorf("Error seeding InviteStatus Table: %v", err)
	}
	utils.Log.Info("InviteStatus Table seed complete.")
	return nil
}
