package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/g3techlabs/revit-api/db"
	"github.com/g3techlabs/revit-api/db/models"
	"gorm.io/gorm"
)

const acceptedStatusId uint = 1
const pendingStatusId uint = 2
const rejectedStatusId uint = 3

type UserRepository interface {
	RegisterUser(user *models.User) error
	FindUserByNickname(nickname string) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	FindUserById(id uint) (*models.User, error)
	UpdateUserPassword(id uint, newPassword string) error
	UpdateUserProfilePic(id uint, newProfilePic *string) error
	Update(id uint, name *string, birthdate *time.Time) error
	GetUsers(page uint, limit uint, nickname string) (*[]models.User, error)
	AreFriends(userId, destintaryId uint) (bool, error)
	RequestFriendship(userId, destinataryId uint) error
	AcceptFriendshipRequest(userId, requesterId uint) error
	RejectFriendshipRequest(userId, requesterId uint) error
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

func (ur *userRepository) UpdateUserProfilePic(id uint, newProfilePic *string) error {
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
	pattern := fmt.Sprintf("%%%s%%", strings.ToLower(nickname))

	query := ur.db.Model(users).Where("nickname LIKE ?", pattern)

	if limit > 0 {
		offset := 0
		if page > 0 {
			offset = int((page - 1) * limit)
		}
		query = query.Limit(int(limit)).Offset(offset)
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	if query.Error != nil {
		return nil, query.Error
	}

	return users, nil
}

func (ur *userRepository) AreFriends(userId, destinataryId uint) (bool, error) {
	var exists bool
	result := ur.db.
		Table("friendship").
		Select("count(*) > 0").
		Where(
			"(requester_id = ? AND receiver_id = ?) OR (requester_id = ? AND receiver_id = ?)",
			userId, destinataryId, destinataryId, userId,
		).
		Where("invite_status_id IN (?, ?)", acceptedStatusId, pendingStatusId).
		Scan(&exists)

	return exists, result.Error
}

func (ur *userRepository) RequestFriendship(userId, destinataryId uint) error {
	friendship := models.Friendship{
		RequesterID:    userId,
		ReceiverID:     destinataryId,
		InviteStatusID: pendingStatusId,
	}

	result := ur.db.Create(&friendship)
	return result.Error
}

func (ur *userRepository) AcceptFriendshipRequest(userId, requesterId uint) error {
	newData := map[string]any{
		"invite_status_id": acceptedStatusId,
		"friends_since":    time.Now().UTC(),
	}

	result := ur.db.
		Model(&models.Friendship{}).
		Where("requester_id = ? AND receiver_id = ? AND invite_status_id = ?", requesterId, userId, pendingStatusId).
		Updates(newData)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("friendship request was not found")
	}

	return result.Error
}

func (ur *userRepository) RejectFriendshipRequest(userId, requesterId uint) error {
	result := ur.db.
		Model(&models.Friendship{}).
		Where("requester_id = ? AND receiver_id = ? AND invite_status_id = ?", requesterId, userId, pendingStatusId).
		Update("invite_status_id", rejectedStatusId)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("friendship request was not found")
	}

	return result.Error
}
