package repository

import (
	"github.com/g3techlabs/revit-api/core/users/models"
	"github.com/g3techlabs/revit-api/db"
	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(user *models.User) error
	FindUserByNickname(nickname string) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	FindUserById(id uint) (*models.User, error)
	UpdateUserPassword(id uint, newPassword string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: db.Db,
	}
}

func (ur userRepository) RegisterUser(user *models.User) error {
	db := db.Db

	result := db.Create(&user)

	return result.Error
}

func (ur userRepository) FindUserByNickname(nickname string) (*models.User, error) {
	db := db.Db
	var user models.User

	result := db.Where("nickname = ?", nickname).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
	}

	return &user, nil
}

func (ur userRepository) FindUserByEmail(email string) (*models.User, error) {
	db := db.Db
	var user models.User

	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
	}

	return &user, nil
}

func (ur userRepository) FindUserById(id uint) (*models.User, error) {
	db := db.Db
	var user models.User

	result := db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
	}

	return &user, nil
}

func (ur userRepository) UpdateUserPassword(id uint, newPassword string) error {
	db := db.Db

	result := db.Table("users").Where("id = ?", id).Update("password", newPassword)

	return result.Error
}
