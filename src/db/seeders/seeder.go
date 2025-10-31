package seeders

import (
	"github.com/g3techlabs/revit-api/src/db"
)

func ExecuteAll() error {
	seedInviteStatusTable(db.Db)
	seedRoleTable(db.Db)
	seedVisibilityTable(db.Db)

	return nil
}
