package repository

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/g3techlabs/revit-api/src/config"
	"github.com/g3techlabs/revit-api/src/core/users/response"
	"github.com/g3techlabs/revit-api/src/db"
	"github.com/g3techlabs/revit-api/src/db/models"
	"gorm.io/gorm"
)

const acceptedStatusId uint = 1
const pendingStatusId uint = 2
const rejectedStatusId uint = 3

var cloudFrontUrl string = config.Get("AWS_CLOUDFRONT_URL")

type UserRepository interface {
	RegisterUser(user *models.User) error
	FindUserByNickname(nickname string) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	FindUserById(id uint) (*models.User, error)
	UpdateUserPassword(id uint, newPassword string) error
	UpdateUserProfilePic(id uint, newProfilePic *string) error
	Update(id uint, name *string, birthdate *time.Time) error
	GetUsers(page uint, limit uint, nickname string) (*[]models.User, error)
	GetFriends(userId, page, limit uint, nickname string) (*[]response.Friend, error)
	AreFriends(userId, destintaryId uint) (bool, error)
	RequestFriendship(userId, destinataryId uint) error
	AcceptFriendshipRequest(userId, requesterId uint) error
	RejectFriendshipRequest(userId, requesterId uint) error
	RemoveFriendship(userId, friendId uint) error
	GetFriendshipRequests(userId uint, page, limit uint) (*response.GetFriendshipRequestsResponse, error)
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

func (ur *userRepository) GetFriends(userId uint, page uint, limit uint, nickname string) (*[]response.Friend, error) {
	friends := new([]response.Friend)
	pattern := fmt.Sprintf("%%%s%%", strings.ToLower(nickname))

	subQuery := ur.db.
		Table("friendship").
		Select(`
		CASE 
			WHEN requester_id = ? THEN receiver_id
			WHEN receiver_id = ? THEN requester_id
		END AS friend_id,
		friends_since
	`, userId, userId).
		Where("(requester_id = ? OR receiver_id = ?)", userId, userId).
		Where("invite_status_id = ?", acceptedStatusId).
		Where("removed_at IS NULL")

	query := ur.db.
		Table("users").
		Select("users.id, users.name, users.nickname, users.email, users.profile_pic, users.created_at AS since, sub.friends_since").
		Joins("JOIN (?) AS sub ON sub.friend_id = users.id", subQuery).
		Where("LOWER(users.nickname) LIKE ?", pattern)

	if limit > 0 {
		offset := 0
		if page > 0 {
			offset = int((page - 1) * limit)
		}
		query = query.Limit(int(limit)).Offset(offset)
	}

	if err := query.Scan(&friends).Error; err != nil {
		return nil, err
	}

	return friends, nil
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
		Where("invite_status_id IN (?, ?) AND removed_at IS NULL", acceptedStatusId, pendingStatusId).
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

func (ur *userRepository) RemoveFriendship(userId, friendId uint) error {
	now := time.Now()

	result := ur.db.Model(&models.Friendship{}).
		Where(
			"(requester_id = ? AND receiver_id = ?) OR (requester_id = ? AND receiver_id = ?)",
			userId, friendId, friendId, userId,
		).
		Where("invite_status_id = ? AND removed_at IS NULL", acceptedStatusId).
		Updates(map[string]interface{}{
			"removed_at":    now,
			"removed_by_id": userId,
		})

	if result.RowsAffected == 0 {
		return fmt.Errorf("friendship was not found")
	}

	return result.Error
}

func (ur *userRepository) GetFriendshipRequests(userId uint, page, limit uint) (*response.GetFriendshipRequestsResponse, error) {
	limitInt := 20
	pageInt := 1
	if limit > 0 {
		limitInt = int(limit)
	}
	if page > 0 {
		pageInt = int(page)
	}

	baseQuery := ur.db.
		Model(&models.Friendship{}).
		Select("friendship.requester_id").
		Where("friendship.invite_status_id = ? AND receiver_id = ? AND removed_at IS NULL AND removed_by_id IS NULL", pendingStatusId, userId)

	// Contar total de registros
	var totalCount int64
	countQuery := ur.db.Raw("SELECT COUNT(*) FROM (?) AS subquery", baseQuery)
	if err := countQuery.Scan(&totalCount).Error; err != nil {
		return nil, err
	}

	// Calcular totalPages
	totalPages := uint(0)
	if totalCount > 0 && limitInt > 0 {
		totalPages = uint(math.Ceil(float64(totalCount) / float64(limitInt)))
	}

	// Buscar os dados paginados
	requests := new([]response.FriendshipRequest)
	query := ur.db.
		Model(&models.Friendship{}).
		Select("requester_id", "requester.nickname", "'"+cloudFrontUrl+"'|| requester.profile_pic AS profile_pic").
		Joins("INNER JOIN users AS requester ON friendship.requester_id = requester.id").
		Where("friendship.invite_status_id = ? AND receiver_id = ? AND removed_at IS NULL AND removed_by_id IS NULL", pendingStatusId, userId)

	if limitInt > 0 {
		offset := (pageInt - 1) * limitInt
		query = query.Limit(limitInt).Offset(offset)
	}

	if err := query.Scan(requests).Error; err != nil {
		return nil, err
	}

	if *requests == nil {
		*requests = make([]response.FriendshipRequest, 0)
	}

	return &response.GetFriendshipRequestsResponse{
		Requests:    *requests,
		CurrentPage: uint(pageInt),
		TotalPages:  totalPages,
	}, nil
}
