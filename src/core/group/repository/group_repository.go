package repository

import (
	"fmt"
	"time"

	"database/sql"

	"github.com/g3techlabs/revit-api/src/config"
	"github.com/g3techlabs/revit-api/src/core/group/input"
	"github.com/g3techlabs/revit-api/src/core/group/response"
	"github.com/g3techlabs/revit-api/src/db"
	"github.com/g3techlabs/revit-api/src/db/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GroupRepository interface {
	CreateGroup(userId uint, data *models.Group) error
	GetGroups(userId uint, filters *input.GetGroupsQuery) (*response.GetGroupsResponse, error)
	GetGroup(userId, groupId uint) (*response.GroupResponse, error)
	GetMembers(groupId uint, queryParams input.GetMembersInput) (*response.GroupMembersResponse, error)
	CanUserViewGroup(userId, groupId uint) (bool, error)
	UpdateMainPhoto(userId, groupId uint, banner string) error
	UpdateBanner(userId, groupId uint, banner string) error
	UpdateGroup(userId, groupId uint, data *input.UpdateGroup) error
	IsUserAdmin(userId, groupId uint) (bool, error)
	InsertNewGroupMember(userId, groupId uint) error
	QuitGroup(userId, groupId uint) error
	MakeGroupInvitation(groupAdminId, groupId, invitedId uint) error
	GetPendingInvites(userId uint, page uint, limit uint) (*[]response.GetPendingInvites, error)
	AcceptPendingInvite(groupId, userId uint) error
	RejectPendingInvite(groupId, userId uint) error
	RemoveMember(groupAdminId, groupId, groupMemberId uint) error
	GetAdminGroups(userId uint, queryParams input.GetAdminGroupsInput) (*response.GetAdminGroupsResponse, error)
}

type groupRepository struct {
	db *gorm.DB
}

func NewGroupRepository() GroupRepository {
	return &groupRepository{
		db: db.Db,
	}
}

var cloudFrontUrl = config.Get("AWS_CLOUDFRONT_URL")

var PublicVisibility uint = 1
var PrivateVisibility uint = 2

var ownerRoleId uint = 1
var adminRoleId uint = 2
var memberRoleId uint = 3

const acceptedStatusId uint = 1
const pendingStatusId uint = 2
const rejectedStatusId uint = 3

func (gr *groupRepository) CreateGroup(userId uint, data *models.Group) error {
	return gr.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(data).Error; err != nil {
			return err
		}

		currentTimestamp := time.Now().UTC()

		ownerData := &models.GroupMember{
			GroupID:        data.ID,
			UserID:         userId,
			RoleID:         ownerRoleId,
			InviteStatusID: acceptedStatusId,
			MemberSince:    &currentTimestamp,
		}

		if err := tx.Create(ownerData).Error; err != nil {
			return err
		}

		return nil
	})
}

func (gr *groupRepository) UpdateMainPhoto(userId, groupId uint, mainPhoto string) error {
	result := gr.db.Model(&models.Group{}).
		Where("id = ? AND EXISTS (SELECT 1 FROM group_member WHERE group_member.group_id = ? AND group_member.user_id = ? AND group_member.invite_status_id = ? AND group_member.role_id IN ?)",
			groupId, groupId, userId, acceptedStatusId, []uint{ownerRoleId, adminRoleId}).
		Update("main_photo", mainPhoto)

	if result.RowsAffected == 0 {
		return fmt.Errorf("group not found or user not allowed")
	}

	return result.Error
}

func (gr *groupRepository) UpdateBanner(userId, groupId uint, banner string) error {
	result := gr.db.Model(&models.Group{}).
		Where("id = ? AND EXISTS (SELECT 1 FROM group_member WHERE group_member.group_id = ? AND group_member.user_id = ? AND group_member.invite_status_id = ? AND group_member.role_id IN ?)",
			groupId, groupId, userId, acceptedStatusId, []uint{ownerRoleId, adminRoleId}).
		Update("banner", banner)

	if result.RowsAffected == 0 {
		return fmt.Errorf("group not found or user not allowed")
	}

	return result.Error
}

func (gr *groupRepository) UpdateGroup(userId, groupId uint, data *input.UpdateGroup) error {
	updates := make(map[string]any)

	if data.Name != nil {
		updates["name"] = *data.Name
	}
	if data.Description != nil {
		updates["description"] = *data.Description
	}
	if data.Visibility != nil {
		switch *data.Visibility {
		case "public":
			updates["visibility_id"] = PublicVisibility
		case "private":
			updates["visibility_id"] = PrivateVisibility
		}
	}
	if data.CityID != nil {
		if err := gr.updateGroupLocation(data.CityID, &updates); err != nil {
			return err
		}
	}

	result := gr.db.Model(&models.Group{}).
		Where("id = ?",
			groupId).
		Where(`EXISTS (SELECT 1 
		FROM group_member 
		WHERE group_member.group_id = ? 
		AND group_member.user_id = ? 
		AND group_member.invite_status_id = ? 
		AND group_member.role_id IN ?)
		`,
			groupId, userId, acceptedStatusId, []uint{ownerRoleId, adminRoleId}).
		Updates(updates)

	if result.RowsAffected == 0 {
		return fmt.Errorf("group not found")
	}

	return result.Error
}

func (gr *groupRepository) updateGroupLocation(cityId *uint, updates *map[string]interface{}) error {
	var newCity sql.NullInt64
	if err := gr.db.Table("city").
		Select("id").
		Where("id = ?", *cityId).
		Scan(&newCity).Error; err != nil {
		return err
	}
	if !newCity.Valid {
		return fmt.Errorf("city not found")
	}

	(*updates)["city_id"] = *cityId

	return nil
}

func (gr *groupRepository) IsUserAdmin(userId, groupId uint) (bool, error) {
	var count int64

	err := gr.db.
		Table("group_member").
		Where("group_id = ? AND user_id = ? AND invite_status_id = ? AND role_id IN ?",
			groupId, userId, acceptedStatusId, []uint{ownerRoleId, adminRoleId}).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (gr *groupRepository) InsertNewGroupMember(userId, groupId uint) error {
	return gr.db.Transaction(func(tx *gorm.DB) error {
		var g models.Group
		if err := tx.Select("id", "visibility_id").Where("id = ?", groupId).First(&g).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return fmt.Errorf("group not found")
			}
			return err
		}

		if g.VisibilityID != PublicVisibility {
			return fmt.Errorf("group is private")
		}

		var count int64
		if err := tx.Model(&models.GroupMember{}).
			Where("group_id = ? AND user_id = ? AND invite_status_id = ? AND left_at IS NULL AND removed_by IS NULL", groupId, userId, acceptedStatusId).
			Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return fmt.Errorf("user already member")
		}

		if err := gr.registerMember(groupId, userId, tx); err != nil {
			return err
		}

		return nil
	})
}

func (gr *groupRepository) registerMember(groupId, userId uint, tx *gorm.DB) error {
	now := time.Now().UTC()
	gm := models.GroupMember{
		GroupID:        groupId,
		UserID:         userId,
		RoleID:         memberRoleId,
		InviteStatusID: acceptedStatusId,
		MemberSince:    &now,
	}

	err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "group_id"}, {Name: "user_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"role_id":          memberRoleId,
			"invite_status_id": acceptedStatusId,
			"member_since":     now,
			"left_at":          nil,
			"removed_by":       nil,
		}),
	}).Create(&gm).Error

	return err
}

func (gr *groupRepository) MakeGroupInvitation(groupAdminId, groupId, invitedId uint) error {
	return gr.db.Transaction(func(tx *gorm.DB) error {
		var admin models.GroupMember
		result := tx.Model(&admin).
			Where("group_id = ? AND user_id = ? AND invite_status_id = ? AND role_id IN ?", groupId, groupAdminId, acceptedStatusId, []uint{ownerRoleId, adminRoleId}).
			Where("left_at IS NULL AND removed_by IS NULL").
			First(&admin)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				return fmt.Errorf("requester not a group admin")
			}
			return result.Error
		}

		var count int64
		result = tx.Model(&models.GroupMember{}).
			Where("group_id = ? AND user_id = ? AND invite_status_id IN ? AND left_at IS NULL AND removed_by IS NULL", groupId, invitedId, []uint{pendingStatusId, acceptedStatusId}).
			Count(&count)
		if result.Error != nil {
			return result.Error
		}

		if count > 0 {
			return fmt.Errorf("invited already invited/a member")
		}

		return gr.makeInvitation(groupAdminId, groupId, invitedId, tx)
	})
}

func (gr *groupRepository) makeInvitation(groupAdminId, groupId, invitedId uint, tx *gorm.DB) error {
	data := models.GroupMember{
		GroupID:        groupId,
		UserID:         invitedId,
		InviterID:      &groupAdminId,
		RoleID:         memberRoleId,
		InviteStatusID: pendingStatusId,
	}

	err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "group_id"}, {Name: "user_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"inviter_id":       groupAdminId,
			"role_id":          memberRoleId,
			"invite_status_id": pendingStatusId,
			"left_at":          nil,
			"removed_by":       nil,
		}),
	}).Create(&data).Error

	return err
}

func (gr *groupRepository) GetPendingInvites(userId uint, page uint, limit uint) (*[]response.GetPendingInvites, error) {
	var pendingInvites []response.GetPendingInvites

	query := gr.db.Model(&models.GroupMember{}).
		Select("g.id as group_id", "g.name AS group_name", "g.main_photo AS group_main_photo", "jsonb_build_object('name', inviter.name, 'inviterProfilePicUrl', '"+cloudFrontUrl+"'|| inviter.profile_pic)").
		Joins("INNER JOIN users AS inviter ON inviter.id = group_member.inviter_id").
		Joins("INNER JOIN groups AS g ON g.id = group_member.group_id").
		Where("user_id = ? AND invite_status_id = ? AND left_at IS NULL AND removed_by IS NULL", userId, pendingStatusId)

	if limit > 0 {
		offset := 0
		if page > 0 {
			offset = int((page - 1) * limit)
		}
		query = query.Limit(int(limit)).Offset(offset)
	}

	if err := query.Scan(&pendingInvites).Error; err != nil {
		return nil, err
	}

	if pendingInvites == nil {
		empty := make([]response.GetPendingInvites, 0)
		return &empty, nil
	}

	return &pendingInvites, nil
}

func (gr *groupRepository) AcceptPendingInvite(groupId uint, userId uint) error {
	data := map[string]interface{}{
		"invite_status_id": acceptedStatusId,
		"member_since":     time.Now().UTC(),
	}

	result := gr.db.
		Model(&models.GroupMember{}).
		Where("user_id = ? AND group_id = ? AND invite_status_id = ?", userId, groupId, pendingStatusId).
		Updates(data)

	if result.RowsAffected == 0 {
		return fmt.Errorf("group invite not found")
	}

	return result.Error
}

func (gr *groupRepository) RejectPendingInvite(groupId, userId uint) error {
	result := gr.db.
		Model(&models.GroupMember{}).
		Where("user_id = ? AND group_id = ? AND invite_status_id = ?", userId, groupId, pendingStatusId).
		Update("invite_status_id", rejectedStatusId)

	if result.RowsAffected == 0 {
		return fmt.Errorf("group invite not found")
	}

	return result.Error
}

func (gr *groupRepository) RemoveMember(groupAdminId, groupId, groupMemberId uint) error {
	data := map[string]interface{}{
		"removed_by": groupAdminId,
		"left_at":    time.Now().UTC(),
	}

	result := gr.db.Model(&models.GroupMember{}).
		Where("user_id = ? AND group_id = ? AND invite_status_id = ? AND left_at IS NULL and removed_by IS NULL", groupMemberId, groupId, acceptedStatusId).
		Updates(data)

	if result.RowsAffected == 0 {
		return fmt.Errorf("member not found")
	}

	return result.Error
}
