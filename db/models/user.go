package models

import (
	"time"
)

type User struct {
	ID         uint
	Name       string `gorm:"not null"`
	Email      string `gorm:"unique;not null"`
	Nickname   string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	ProfilePic *string
	CreatedAt  time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt  time.Time `gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}
