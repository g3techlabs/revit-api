package models

type State struct {
	ID        uint
	Name      string `gorm:"not null"`
	Acronym   string `gorm:"not null"`
	CountryID uint   `gorm:"not null"`

	Cities []City `gorm:"foreignKey:StateID;constraint:OnDelete:CASCADE;"`
}
