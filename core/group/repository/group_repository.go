package repository

import (
	"fmt"

	"github.com/g3techlabs/revit-api/db"
	"github.com/g3techlabs/revit-api/db/models"
	"gorm.io/gorm"
)

type GroupRepository interface {
	CreateGroup(userId uint, data *models.Group) error
	UpdateMainPhoto(userId, groupId uint, banner string) error
	UpdateBanner(userId, groupId uint, banner string) error
}

type groupRepository struct {
	db *gorm.DB
}

func NewGroupRepository() GroupRepository {
	return &groupRepository{
		db: db.Db,
	}
}

var PublicVisibility uint = 1
var PrivateVisibility uint = 2

var ownerId uint = 1
var adminId uint = 2
var memberId uint = 3

const acceptedStatusId uint = 1
const pendingStatusId uint = 2
const rejectedStatusId uint = 3

func (gr *groupRepository) CreateGroup(userId uint, data *models.Group) error {
	return gr.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(data).Error; err != nil {
			return err
		}

		ownerData := &models.GroupMember{
			GroupID:        data.ID,
			UserID:         userId,
			RoleID:         ownerId,
			InviteStatusID: acceptedStatusId,
		}

		if err := tx.Create(ownerData).Error; err != nil {
			return err
		}

		return nil
	})
}

func (gr *groupRepository) UpdateMainPhoto(userId, groupId uint, mainPhoto string) error {
	result := gr.db.Model(&models.Group{}).
		Where("id = ? AND EXISTS (SELECT 1 FROM group_members WHERE group_members.group_id = ? AND group_members.user_id = ? AND group_members.invite_status_id = ? AND group_members.role_id IN ?)",
			groupId, groupId, userId, acceptedStatusId, []uint{ownerId, adminId}).
		Update("main_photo", mainPhoto)

	if result.RowsAffected == 0 {
		return fmt.Errorf("group not found or user not allowed")
	}

	return result.Error
}

func (gr *groupRepository) UpdateBanner(userId, groupId uint, banner string) error {
	result := gr.db.Model(&models.Group{}).
		Where("id = ? AND EXISTS (SELECT 1 FROM group_members WHERE group_members.group_id = ? AND group_members.user_id = ? AND group_members.invite_status_id = ? AND group_members.role_id IN ?)",
			groupId, groupId, userId, acceptedStatusId, []uint{ownerId, adminId}).
		Update("banner", banner)

	if result.RowsAffected == 0 {
		return fmt.Errorf("group not found or user not allowed")
	}

	return result.Error
}
