package models

import "time"

type GroupMember struct {
	ID             uint
	MemberSince    *time.Time
	GroupID        uint `gorm:"not null"`
	UserID         uint `gorm:"not null"`
	InviterID      *uint
	RoleID         uint `gorm:"not null"`
	InviteStatusID uint `gorm:"not null"`

	Group   Group `gorm:"foreignKey:GroupID"`
	User    User  `gorm:"foreignKey:UserID"`
	Inviter *User `gorm:"foreignKey:InviterID"`
}
