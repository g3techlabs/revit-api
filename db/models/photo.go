package models

import "time"

type Photo struct {
	ID        uint
	Reference string `gorm:"not null"`
	DeletedAt *time.Time
	VehicleID uint `gorm:"not null"`
}
