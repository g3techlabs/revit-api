package models

import (
	"time"
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
