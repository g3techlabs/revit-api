package extensions

import (
	"github.com/g3techlabs/revit-api/src/utils"
	"gorm.io/gorm"
)

func EnablePostGISExtension(db *gorm.DB) {
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS postgis;`).Error; err != nil {
		utils.Log.Fatalf("Error creating PostGIS extension: %v", err)
	}
}
