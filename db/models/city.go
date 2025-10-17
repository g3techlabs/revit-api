package models

type City struct {
	ID       uint
	Name     string `gorm:"not null"`
	Location string `gorm:"type:GEOGRAPHY(POINT,4326);not null"`
	StateID  uint   `gorm:"not null"`

	Groups []Group `gorm:"foreignKey:CityID;constraint:OnDelete:SET NULL;"`
}
