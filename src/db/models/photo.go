package models

import (
	"time"

	"github.com/g3techlabs/revit-api/src/core/vehicle/response"
)

type Photo struct {
	ID        uint
	Reference string `gorm:"not null"`
	DeletedAt *time.Time
	VehicleID uint `gorm:"not null"`
}

func (p *Photo) ToPhotoResponse() *response.Photo {
	return &response.Photo{
		ID:  p.ID,
		Url: p.Reference,
	}
}
