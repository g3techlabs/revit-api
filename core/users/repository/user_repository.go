package repository

import (
	"fmt"
	"time"

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
	UpdateUserProfilePic(id uint, newProfilePic string) error
	Update(id uint, name *string, birthdate *time.Time) error
	GetUsers(page uint, limit uint, nickname string) (*[]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: db.Db,
	}
}

func (ur *userRepository) RegisterUser(user *models.User) error {
	result := ur.db.Create(&user)

	return result.Error
}

func (ur *userRepository) FindUserByNickname(nickname string) (*models.User, error) {
	var user models.User

	result := ur.db.Where("nickname = ?", nickname).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
	}

	return &user, nil
}

func (ur *userRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User

	result := ur.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
	}

	return &user, nil
}

func (ur *userRepository) FindUserById(id uint) (*models.User, error) {
	var user models.User

	result := ur.db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
	}

	return &user, nil
}

func (ur *userRepository) UpdateUserPassword(id uint, newPassword string) error {
	result := ur.db.Table("users").Where("id = ?", id).Update("password", newPassword)

	return result.Error
}

func (ur *userRepository) UpdateUserProfilePic(id uint, newProfilePic string) error {
	result := ur.db.Table("users").Where("id = ?", id).Update("profile_pic", newProfilePic)

	return result.Error
}

func (ur *userRepository) Update(id uint, name *string, birthdate *time.Time) error {
	data := map[string]interface{}{}

	if name != nil {
		data["name"] = *name
	}
	if birthdate != nil {
		data["birthdate"] = *birthdate
	}

	if len(data) == 0 {
		return nil
	}

	result := ur.db.Table("users").Where("id = ?", id).Updates(data)

	return result.Error
}

func (ur *userRepository) GetUsers(page uint, limit uint, nickname string) (*[]models.User, error) {
	users := new([]models.User)

	pattern := fmt.Sprintf("%%%s%%", nickname)

	query := ur.db.Model(users).Where("nickname LIKE ?", pattern).Limit(int(limit)).Offset(int((page - 1))).Find(&users)

	if query.Error != nil {
		return nil, query.Error
	}

	return users, nil

}
