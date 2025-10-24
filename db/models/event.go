package models

import "time"

type Event struct {
	ID           uint
	Name         string    `gorm:"not null"`
	Description  string    `gorm:"not null"`
	Date         time.Time `gorm:"not null"`
	Photo        *string
	Location     string `gorm:"type:GEOGRAPHY(POINT,4326);not null"`
	Canceled     bool   `gorm:"not null;default:false"`
	City         string `gorm:"not null"`
	VisibilityID uint   `gorm:"not null"`
	GroupID      *uint
	CreatedAt    time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt    time.Time `gorm:"not null"`

	Visibility Visibility `gorm:"foreignKey:VisibilityID;references:ID;constraint:OnDelete:SET NULL"`
	Group      *Group     `gorm:"foreignKey:GroupID;references:ID;constraint:OnDelete:SET NULL"`
}
