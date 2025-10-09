package models

import (
	"time"

	"github.com/g3techlabs/revit-api/core/users/response"
)

type User struct {
	ID         uint
	Name       string `gorm:"not null"`
	Email      string `gorm:"unique;not null"`
	Nickname   string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	ProfilePic *string
	Birthdate  *time.Time `gorm:"type:date"`
	CreatedAt  time.Time  `gorm:"not null;default:current_timestamp"`
	UpdatedAt  time.Time  `gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}

func (u User) ToUserCreatedResponse() *response.UserCreatedResponse {
	return &response.UserCreatedResponse{
		Name:       u.Name,
		Email:      u.Email,
		Nickname:   u.Nickname,
		ProfilePic: u.ProfilePic,
		Birthdate:  u.Birthdate,
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}

func (u User) ToGetUserResponse() *response.GetUserResponse {
	return &response.GetUserResponse{
		ID:            u.ID,
		Name:          u.Name,
		Nickname:      u.Nickname,
		ProfilePicUrl: u.ProfilePic,
		Since:         &u.CreatedAt,
	}
}
