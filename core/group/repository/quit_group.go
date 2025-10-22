package repository

import (
	"fmt"
	"time"

	"github.com/g3techlabs/revit-api/db/models"
	"gorm.io/gorm"
)

func (gr *groupRepository) QuitGroup(userId, groupId uint) error {
	return gr.db.Transaction(func(tx *gorm.DB) error {

		var groupMember models.GroupMember
		if err := gr.queryQuitter(userId, groupId, &groupMember, tx); err != nil {
			return err
		}

		if err := gr.markQuitterLeftAt(userId, groupId, tx); err != nil {
			return err
		}

		if err := gr.makeNewGroupOwnerIfNeeded(groupMember.RoleID == ownerRoleId, groupId, tx); err != nil {
			return err
		}

		return nil
	})
}

func (gr *groupRepository) queryQuitter(userId, groupId uint, model *models.GroupMember, tx *gorm.DB) error {
	if err := tx.Select("role_id").
		Where("group_id = ? AND user_id = ? AND invite_status_id = ? AND left_at IS NULL AND removed_by IS NULL", groupId, userId, acceptedStatusId).
		Order("member_since ASC").
		First(model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("user is not a member")
		}
		return err
	}

	return nil
}

func (gr *groupRepository) markQuitterLeftAt(userId, groupId uint, tx *gorm.DB) error {
	result := tx.
		Table("group_member").
		Where("group_id = ? AND user_id = ? AND invite_status_id = ? AND left_at IS NULL AND removed_by IS NULL", groupId, userId, acceptedStatusId).
		Update("left_at", time.Now().UTC())

	if result.RowsAffected == 0 {
		return fmt.Errorf("user is not a member")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (gr *groupRepository) makeNewGroupOwnerIfNeeded(isQuitterTheOwner bool, groupId uint, tx *gorm.DB) error {
	if !isQuitterTheOwner {
		return nil
	}

	var ownerCandidate models.GroupMember
	if err := gr.queryNewOwnerCandidate(groupId, tx, &ownerCandidate); err != nil {
		return err
	}

	if err := gr.makeNewOwner(&ownerCandidate, tx); err != nil {
		return err
	}

	return nil
}

func (gr *groupRepository) queryNewOwnerCandidate(groupId uint, tx *gorm.DB, model *models.GroupMember) error {
	if err := gr.queryMember(groupId, true, model, tx); err != nil {
		return err
	}

	if model.UserID != 0 {
		return nil
	}

	if err := gr.queryMember(groupId, false, model, tx); err != nil {
		return err
	}

	return nil
}

func (gr *groupRepository) queryMember(groupId uint, admin bool, model *models.GroupMember, tx *gorm.DB) error {
	query := tx.Where("group_id = ? AND invite_status_id = ? AND left_at IS NULL AND removed_by IS NULL", groupId, acceptedStatusId)

	if admin {
		query = query.Where("role_id = ?", adminRoleId)
	}

	if err := query.Order("member_since ASC").First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}

	return nil
}

func (gr *groupRepository) makeNewOwner(model *models.GroupMember, tx *gorm.DB) error {
	result := tx.Model(model).Update("role_id", ownerRoleId)

	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return result.Error
}
