package models

import "time"

type RouteParticipant struct {
	UserID        uint `gorm:"primaryKey;autoIncrement:false"`
	RouteID       uint `gorm:"primaryKey;autoIncrement:false"`
	StartLocation any  `gorm:"type:GEOGRAPHY(POINT,4326);not null"`
	IsOwner       bool `gorm:"not null;default:false"`
	EndedAt       *time.Time

	User  User  `gorm:"foreignKey:UserID"`
	Route Route `gorm:"foreignKey:RouteID"`
}
