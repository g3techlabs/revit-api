package models

import "time"

type Group struct {
	ID           uint
	Name         string `gorm:"not null"`
	Description  string `gorm:"not null"`
	Photo        *string
	Banner       *string
	CreatedAt    time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt    time.Time `gorm:"not null"`
	VisibilityID uint      `gorm:"not null"`
	CityID       uint      `gorm:"not null"`

	Members []GroupMember `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE;"`
}
