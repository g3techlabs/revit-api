package models

import "time"

type GroupMember struct {
	GroupID        uint `gorm:"primaryKey;autoIncrement:false"`
	UserID         uint `gorm:"primaryKey;autoIncrement:false"`
	MemberSince    *time.Time
	InviterID      *uint
	RoleID         uint `gorm:"not null"`
	InviteStatusID uint `gorm:"not null"`
	LeftAt         *time.Time
	RemovedBy      *uint

	Group         Group `gorm:"foreignKey:GroupID"`
	User          User  `gorm:"foreignKey:UserID"`
	Inviter       *User `gorm:"foreignKey:InviterID"`
	RemovedByUser *User `gorm:"foreignKey:RemovedBy"`
}
