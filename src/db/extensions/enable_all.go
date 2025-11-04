package extensions

import (
	"github.com/g3techlabs/revit-api/src/db"
	"github.com/g3techlabs/revit-api/src/utils"
)

func EnableAll(logger utils.ILogger) {
	EnablePostGISExtension(db.Db, logger)
	EnablePostTRGMExtension(db.Db, logger)
}
