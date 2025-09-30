package repository

import (
	"github.com/g3techlabs/revit-api/db"
	"github.com/g3techlabs/revit-api/db/models"
	"gorm.io/gorm"
)

type User = models.User

func RegisterUser(user *User) error {
	db := db.Db

	result := db.Create(&user)

	return result.Error
}

func FindUserByNickname(nickname string) (*User, error) {
	db := db.Db
	var user User

	result := db.Where("nickname = ?", nickname).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
	}

	return &user, nil
}

func FindUserByEmail(email string) (*User, error) {
	db := db.Db
	var user User

	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
	}

	return &user, nil
}

func FindUserById(id uint) (*User, error) {
	db := db.Db
	var user User

	result := db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
	}

	return &user, nil
}

func UpdateUserPassword(id uint, newPassword string) error {
	db := db.Db

	result := db.Table("users").Where("id = ?", id).Update("password", newPassword)

	return result.Error
}
