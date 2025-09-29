package models

import "time"

type ResetToken struct {
	ID        uint
	UserID    uint `gorm:"not null;unique"`
	User      User
	Token     string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;default:current_timestamp"`
}
