package models

import "time"

type Friendship struct {
	RequesterID    uint `gorm:"primaryKey"`
	ReceiverID     uint `gorm:"primaryKey"`
	FriendsSince   *time.Time
	InviteStatusID uint `gorm:"not null"`
	RemovedAt      *time.Time
	RemovedByID    *uint

	RequesterIDRef uint         `gorm:"column:requester_id"`
	ReceiverIDRef  uint         `gorm:"column:receiver_id"`
	InviteStatus   InviteStatus `gorm:"foreignKey:InviteStatusID;references:ID;constraint:OnDelete:CASCADE;not null"`

	// For preventing infinite recursion
	Requester *User `gorm:"foreignKey:RequesterID;references:ID;-" json:"-"`
	Receiver  *User `gorm:"foreignKey:ReceiverID;references:ID;-" json:"-"`
	RemovedBy *User `gorm:"foreignKey:RemovedByID;references:ID;constraint:OnDelete:SET NULL;" json:"-"`
}
