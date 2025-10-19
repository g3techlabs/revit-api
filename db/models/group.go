package models

import "time"

type Group struct {
	ID           uint
	Name         string `gorm:"not null"`
	Description  string `gorm:"not null"`
	MainPhoto    *string
	Banner       *string
	CreatedAt    time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt    time.Time `gorm:"not null"`
	VisibilityID uint      `gorm:"not null"`
	CityID       uint      `gorm:"not null"`

	Visibility Visibility    `gorm:"foreignKey:VisibilityID;references:ID;constraint:OnDelete:SET NULL"`
	City       City          `gorm:"foreignKey:CityID;references:ID;constraint:OnDelete:SET NULL"`
	Members    []GroupMember `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE;"`
}

func (Group) TableName() string {
	return "groups"
}
