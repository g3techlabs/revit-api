package extensions

import (
	"github.com/g3techlabs/revit-api/db"
)

func EnableAll() {
	EnablePostGISExtension(db.Db)
}
