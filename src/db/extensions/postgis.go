package extensions

import (
	"github.com/g3techlabs/revit-api/src/utils"
	"gorm.io/gorm"
)

func EnablePostGISExtension(db *gorm.DB, logger utils.ILogger) {
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS postgis;`).Error; err != nil {
		logger.Fatalf("Error creating PostGIS extension: %v", err)
	}
}
