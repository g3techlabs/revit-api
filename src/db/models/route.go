package models

import "time"

type Route struct {
	ID          uint
	IsDone      bool `gorm:"not null;default:false"`
	Destination any  `gorm:"type:GEOGRAPHY(POINT,4326);not null"`
	StartedAt   *time.Time
	FinishedAt  *time.Time

	Participants []RouteParticipant `gorm:"foreignKey:RouteID;constraint:OnDelete:CASCADE;"`
}
