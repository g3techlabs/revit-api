package models

type Role struct {
	ID   uint
	Name string `gorm:"not null;unique"`
}
