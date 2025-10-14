package repository

import (
	"fmt"
	"strings"

	"github.com/g3techlabs/revit-api/db"
	"github.com/g3techlabs/revit-api/db/models"
	"gorm.io/gorm"
)

type VehicleRepository interface {
	CreateVehicle(data *models.Vehicle) error
	UpdateMainPhoto(vehicleId uint, mainPhotoKey string) error
	UpdateVehicleInfo(vehicleId uint, data *models.Vehicle) error
	GetVehicles(userId, page, limit uint, nickname string) (*[]models.Vehicle, error)
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
	vr.lowerStrings(data)

	result := vr.db.Create(data)

	return result.Error
}

func (vr *vehicleRepository) GetVehicles(userId, page, limit uint, nickname string) (*[]models.Vehicle, error) {
	vehicles := new([]models.Vehicle)
	pattern := fmt.Sprintf("%%%s%%", strings.ToLower(nickname))

	query := vr.db.
		Preload("Photos").
		Where("user_id = ? AND deleted_at IS NULL", userId).
		Where("nickname LIKE ?", pattern).
		Order("created_at DESC")

	if limit > 0 {
		offset := 0
		if page > 0 {
			offset = int((page - 1) * limit)
		}
		query = query.Limit(int(limit)).Offset(offset)
	}

	if err := query.Find(&vehicles).Error; err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (vr *vehicleRepository) UpdateMainPhoto(vehicleId uint, mainPhotoKey string) error {
	result := vr.db.Table("vehicle").Where("id = ?", vehicleId).Update("main_photo", mainPhotoKey)

	if result.RowsAffected == 0 {
		return fmt.Errorf("vehicle not found")
	}

	return result.Error
}

func (vr *vehicleRepository) UpdateVehicleInfo(vehicleId uint, data *models.Vehicle) error {
	vr.lowerStrings(data)

	result := vr.db.Model(&models.Vehicle{}).
		Where("id = ?", vehicleId).
		Updates(data)

	if result.RowsAffected == 0 {
		return fmt.Errorf("vehicle not found")
	}

	return result.Error
}

func (*vehicleRepository) lowerStrings(i *models.Vehicle) {
	i.Nickname = strings.ToLower(i.Nickname)
	i.Brand = strings.ToLower(i.Brand)
	i.Model = strings.ToLower(i.Model)
	if i.Version != nil {
		lowered := strings.ToLower(*i.Version)
		i.Version = &lowered
	}
}
