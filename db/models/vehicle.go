package models

import (
	"time"

	"github.com/g3techlabs/revit-api/core/vehicle/response"
)

type Vehicle struct {
	ID        uint
	Nickname  string `gorm:"not null"`
	Brand     string `gorm:"not null"`
	Model     string `gorm:"not null"`
	Year      uint   `gorm:"not null"`
	Version   *string
	MainPhoto *string
	CreatedAt time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt *time.Time
	UserID    uint `gorm:"not null"`

	Photos []Photo `gorm:"constraint:OnDelete:CASCADE"`
}

func (v *Vehicle) ToVehicleResponse() *response.Vehicle {
	return &response.Vehicle{
		ID:           v.ID,
		Nickname:     v.Nickname,
		Brand:        v.Brand,
		Model:        v.Model,
		Year:         v.Year,
		MainPhotoUrl: v.MainPhoto,
		CreatedAt:    v.CreatedAt,
	}
}
