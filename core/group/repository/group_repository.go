package repository

import (
	"fmt"
	"strings"
	"time"

	"database/sql"

	"github.com/g3techlabs/revit-api/config"
	"github.com/g3techlabs/revit-api/core/group/input"
	"github.com/g3techlabs/revit-api/core/group/response"
	"github.com/g3techlabs/revit-api/db"
	"github.com/g3techlabs/revit-api/db/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GroupRepository interface {
	CreateGroup(userId uint, data *models.Group) error
	GetGroups(userId uint, filters *input.GetGroupsQuery) (*[]response.GetGroupsResponse, error)
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

func (gr *groupRepository) GetGroups(userId uint, filters *input.GetGroupsQuery) (*[]response.GetGroupsResponse, error) {

	groups := new([]response.GetGroupsResponse)

	query := `
		SELECT 
			g.id,
			g.name,
			g.description,
			CAST(@cloudFrontUrl AS text) || g.main_photo AS "main_photo",
			CAST(@cloudFrontUrl AS text) || g.banner AS "banner",
			g.created_at,
			v.name AS "visibility",
			c.name AS "city",
			s.name AS "state",
			r.name AS "member_type",
			COALESCE(
				json_agg(
					DISTINCT jsonb_build_object(
						'name', u.name,
						'profilePicUrl', u.profile_pic_url
					)
				) FILTER (WHERE u.user_id IS NOT NULL),
				'[]'
			) AS "friends_in_group"
		FROM groups g
		JOIN visibility v ON v.id = g.visibility_id
		JOIN city c ON c.id = g.city_id
		JOIN state s ON s.id = c.state_id
		LEFT JOIN group_member gm_user ON gm_user.group_id = g.id AND gm_user.user_id = @userId AND invite_status_id = @acceptedStatusId
		LEFT JOIN role r ON r.id = gm_user.role_id
		LEFT JOIN LATERAL (
			SELECT 
				gm_friend.user_id,
				users.name,
				@cloudFrontUrl || users.profile_pic as profile_pic_url
			FROM group_member gm_friend
			JOIN users ON users.id = gm_friend.user_id
			JOIN friendship f ON (
				(f.requester_id = @userId AND f.receiver_id = users.id)
				OR (f.receiver_id = @userId AND f.requester_id = users.id)
			)
			WHERE gm_friend.group_id = g.id
			ORDER BY users.name
			LIMIT 3
		) AS u ON TRUE
	`

	params := []interface{}{
		sql.Named("userId", userId),
		sql.Named("cloudFrontUrl", cloudFrontUrl),
		sql.Named("acceptedStatusId", acceptedStatusId),
	}

	gr.buildGetGroupsWhereStatement(&query, &params, filters)

	query += `
		GROUP BY g.id, v.name, c.name, s.name, r.name
		ORDER BY g.created_at DESC
	`
	gr.addPagination(&query, &params, filters.Limit, filters.Page)

	if err := gr.db.Raw(query, params...).Scan(groups).Error; err != nil {
		return nil, err
	}

	if *groups == nil {
		empty := make([]response.GetGroupsResponse, 0)
		return &empty, nil
	}

	return groups, nil
}

func (gr *groupRepository) buildGetGroupsWhereStatement(query *string, params *[]interface{}, filters *input.GetGroupsQuery) {
	conditions := []string{}

	if filters.Name != "" {
		conditions = append(conditions, "LOWER(g.name) LIKE LOWER(@name)")
		*params = append(*params, sql.Named("name", "%"+filters.Name+"%"))
	}
	if filters.CityId != 0 {
		conditions = append(conditions, "g.city_id = @cityId")
		*params = append(*params, sql.Named("cityId", filters.CityId))
	}
	if filters.StateId != 0 {
		conditions = append(conditions, "c.state_id = @stateId")
		*params = append(*params, sql.Named("stateId", filters.StateId))
	}
	if filters.Visibility != "" {
		conditions = append(conditions, "LOWER(v.name) = LOWER(@visibility)")
		*params = append(*params, sql.Named("visibility", filters.Visibility))
	}
	if filters.Member {
		conditions = append(conditions, "gm_user.user_id IS NOT NULL")
	}

	groupHasActiveUsersFilter := "EXISTS (SELECT 1 FROM group_member WHERE group_member.group_id = g.id AND left_at IS NULL)"
	groupIsPublicOrUserIsAMemberFilter := "(g.visibility_id = @publicVisibility OR gm_user.user_id IS NOT NULL)"
	conditions = append(conditions, groupHasActiveUsersFilter, groupIsPublicOrUserIsAMemberFilter)
	*params = append(*params, sql.Named("publicVisibility", PublicVisibility))

	if len(conditions) > 0 {
		*query += " WHERE " + strings.Join(conditions, " AND ")
	}
}

func (gr *groupRepository) addPagination(query *string, params *[]interface{}, queryLimit uint, queryPage uint) {
	limit := 20
	page := 1
	if queryLimit > 0 {
		limit = int(queryLimit)
	}
	if queryPage > 0 {
		page = int(queryPage)
	}
	offset := (page - 1) * limit

	*query += " LIMIT @limit OFFSET @offset"

	*params = append(*params, sql.Named("limit", limit), sql.Named("offset", offset))
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
	result := gr.db.
		Model(&models.GroupMember{}).
		Where("user_id = ? AND group_id = ? AND invite_status_id = ?", userId, groupId, pendingStatusId).
		Update("invite_status_id", acceptedStatusId)

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
