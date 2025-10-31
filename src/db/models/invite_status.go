package models

type InviteStatus struct {
	ID     uint
	Status string `gorm:"unique;not null"`
}
