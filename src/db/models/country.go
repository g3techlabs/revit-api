package models

type Country struct {
	ID      uint
	Name    string `gorm:"not null"`
	Acronym string `gorm:"not null"`

	States []State `gorm:"foreignKey:CountryID;constraint:OnDelete:CASCADE;"`
}
