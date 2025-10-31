package repository

import (
	"fmt"
	"strings"

	"github.com/g3techlabs/revit-api/src/db"
	"github.com/g3techlabs/revit-api/src/db/models"
	"gorm.io/gorm"
)

type VehicleRepository interface {
	CreateVehicle(data *models.Vehicle) error
	GetVehicles(userId, page, limit uint, nickname string) (*[]models.Vehicle, error)
	InsertPhoto(vehicleId uint, photoReference string) error
	UpdateMainPhoto(vehicleId uint, mainPhotoKey string) error
	UpdateVehicleInfo(vehicleId uint, data *models.Vehicle) error
	MarkPhotoAsRemoved(userId, vehicleId, photoId uint) error
	MarkVehicleAsRemoved(userId, vehicleId uint) error
	DeleteMainPhoto(userId, vehicleId uint) error
	IsVehicleAvailable(userId, vehicleId uint) (bool, error)
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
		Preload("Photos", "deleted_at IS NULL").
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

func (vr *vehicleRepository) InsertPhoto(vehicleId uint, photoReference string) error {
	data := &models.Photo{
		Reference: photoReference,
		VehicleID: vehicleId,
	}

	result := vr.db.Create(data)

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

func (vr *vehicleRepository) MarkVehicleAsRemoved(userId, vehicleId uint) error {
	result := vr.db.Model(&models.Vehicle{}).
		Where("id = ? AND user_id = ? AND deleted_at IS NULL", vehicleId, userId).
		Update("deleted_at", gorm.Expr("NOW()"))

	if result.RowsAffected == 0 {
		return fmt.Errorf("vehicle not found")
	}

	return result.Error
}

func (vr *vehicleRepository) DeleteMainPhoto(userId, vehicleId uint) error {
	result := vr.db.Model(&models.Vehicle{}).
		Where("id = ? AND user_id = ? AND deleted_at IS NULL", vehicleId, userId).
		Update("main_photo", nil)

	if result.RowsAffected == 0 {
		return fmt.Errorf("vehicle not found")
	}

	return result.Error
}

func (vr *vehicleRepository) MarkPhotoAsRemoved(userId, vehicleId, photoId uint) error {
	result := vr.db.Model(&models.Photo{}).
		Where("id = ? AND vehicle_id = ? AND vehicle_id IN (SELECT id FROM vehicle WHERE id = ? AND user_id = ? AND deleted_at IS NULL)", photoId, vehicleId, vehicleId, userId).
		Where("deleted_at IS NULL").
		Update("deleted_at", gorm.Expr("NOW()"))

	if result.RowsAffected == 0 {
		return fmt.Errorf("photo not found")
	}

	return result.Error
}

func (vr *vehicleRepository) IsVehicleAvailable(userId, vehicleId uint) (bool, error) {
	var count int64
	err := vr.db.Model(&models.Vehicle{}).
		Where("id = ? AND user_id = ? AND deleted_at IS NULL", vehicleId, userId).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
