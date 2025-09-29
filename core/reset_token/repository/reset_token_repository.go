package repository

import (
	"github.com/g3techlabs/revit-api/db"
	"github.com/g3techlabs/revit-api/db/models"
)

func RegisterResetToken(resetToken *models.ResetToken) error {
	db := db.Db

	result := db.Create(&resetToken)

	return result.Error
}
