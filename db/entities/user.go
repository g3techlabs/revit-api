package entities

import (
	"time"
)

type User struct {
	ID         uint
	Name       string
	Email      string
	Nickname   string
	Password   string
	ProfilePic string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (User) TableName() string {
	return "users"
}
