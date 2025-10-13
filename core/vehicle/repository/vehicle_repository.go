package repository

import (
	"fmt"

	"github.com/g3techlabs/revit-api/db"
	"github.com/g3techlabs/revit-api/db/models"
	"gorm.io/gorm"
)

type VehicleRepository interface {
	CreateVehicle(data *models.Vehicle) error
	UpdateMainPhoto(vehicleId uint, mainPhotoKey string) error
}

type vehicleRepository struct {
	db *gorm.DB
}

func NewVehicleRepository() VehicleRepository {
	return &vehicleRepository{
		db: db.Db,
	}
}

func (vr *vehicleRepository) CreateVehicle(data *models.Vehicle) error {
	result := vr.db.Create(data)

	return result.Error
}

func (vr *vehicleRepository) UpdateMainPhoto(vehicleId uint, mainPhotoKey string) error {
	result := vr.db.Table("vehicle").Where("id = ?", vehicleId).Update("main_photo", mainPhotoKey)

	if result.RowsAffected == 0 {
		return fmt.Errorf("vehicle not found")
	}

	return result.Error
}
