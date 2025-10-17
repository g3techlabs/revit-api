package models

type Visibility struct {
	ID   uint
	Name string `gorm:"not null;unique"`
}
