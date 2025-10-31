package extensions

import (
	"github.com/g3techlabs/revit-api/src/utils"
	"gorm.io/gorm"
)

func EnablePostTRGMExtension(db *gorm.DB) {
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS pg_trgm;`).Error; err != nil {
		utils.Log.Fatalf("Error creating pg_trgm extension: %v", err)
	}
}
