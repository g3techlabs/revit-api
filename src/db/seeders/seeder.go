package seeders

import (
	"github.com/g3techlabs/revit-api/src/db"
	"github.com/g3techlabs/revit-api/src/utils"
)

func ExecuteAll(logger utils.ILogger) error {
	seedInviteStatusTable(db.Db, logger)
	seedRoleTable(db.Db, logger)
	seedVisibilityTable(db.Db, logger)

	return nil
}
