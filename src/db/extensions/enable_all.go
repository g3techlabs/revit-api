package extensions

import (
	"github.com/g3techlabs/revit-api/src/db"
)

func EnableAll() {
	EnablePostGISExtension(db.Db)
	EnablePostTRGMExtension(db.Db)
}
