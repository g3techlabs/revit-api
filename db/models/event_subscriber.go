package models

import "time"

type EventSubscriber struct {
	EventID        uint `gorm:"primaryKey;autoIncrement:false"`
	UserID         uint `gorm:"primaryKey;autoIncrement:false"`
	InviteStatusID uint `gorm:"not null"`
	RoleID         uint `gorm:"not null"`
	InviterID      *uint
	LeftAt         *time.Time
	RemovedBy      *uint

	User          User         `gorm:"foreignKey:UserID"`
	Event         Event        `gorm:"foreignKey:EventID"`
	InviteStatus  InviteStatus `gorm:"foreignKey:InviteStatusID;references:ID;constraint:OnDelete:CASCADE;not null"`
	Role          Role         `gorm:"foreignKey:RoleID;references:ID;constraint:OnDelete:CASCADE;not null"`
	Inviter       *User        `gorm:"foreignKey:InviterID;references:ID;constraint:OnDelete:SET NULL;"`
	RemovedByUser *User        `gorm:"foreignKey:RemovedBy;references:ID;constraint:OnDelete:SET NULL;"`
}
