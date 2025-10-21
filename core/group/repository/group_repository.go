package repository

import (
	"fmt"
	"strings"

	"database/sql"

	"github.com/g3techlabs/revit-api/config"
	"github.com/g3techlabs/revit-api/core/group/input"
	"github.com/g3techlabs/revit-api/core/group/response"
	"github.com/g3techlabs/revit-api/db"
	"github.com/g3techlabs/revit-api/db/models"
	"gorm.io/gorm"
)

type GroupRepository interface {
	CreateGroup(userId uint, data *models.Group) error
	GetGroups(userId uint, filters *input.GetGroupsQuery) (*[]response.GetGroupsResponse, error)
	UpdateMainPhoto(userId, groupId uint, banner string) error
	UpdateBanner(userId, groupId uint, banner string) error
	UpdateGroup(userId, groupId uint, data *input.UpdateGroup) error
	IsUserAdmin(userId, groupId uint) (bool, error)
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

func (gr *groupRepository) GetGroups(userId uint, filters *input.GetGroupsQuery) (*[]response.GetGroupsResponse, error) {
	cloudFrontUrl := config.Get("AWS_CLOUDFRONT_URL")

	groups := new([]response.GetGroupsResponse)

	query := `
		SELECT 
			g.id,
			g.name,
			g.description,
			@cloudFrontUrl || g.main_photo AS "mainPhoto",
			@cloudFrontUrl || g.banner AS "banner",
			g.created_at,
			v.name AS "visibility",
			c.name AS "city",
			s.name AS "state",
			r.name AS "memberType",
			COALESCE(
				json_agg(
					DISTINCT jsonb_build_object(
						'name', u.name,
						'profilePicUrl', u.profile_pic_url
					)
				) FILTER (WHERE u.user_id IS NOT NULL),
				'[]'
			) AS "friendsInGroup"
		FROM groups g
		JOIN visibility v ON v.id = g.visibility_id
		JOIN city c ON c.id = g.city_id
		JOIN state s ON s.id = c.state_id
		LEFT JOIN group_member gm_user ON gm_user.group_id = g.id AND gm_user.user_id = @userId
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
	if data.CityID != nil {
		updates["city_id"] = *data.CityID
	}
	if data.Visibility != nil {
		switch *data.Visibility {
		case "public":
			updates["visibility_id"] = PublicVisibility
		case "private":
			updates["visibility_id"] = PrivateVisibility
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
		AND group_member.role_id IN ?)`,
			groupId, userId, acceptedStatusId, []uint{ownerId, adminId}).
		Updates(updates)

	if result.RowsAffected == 0 {
		return fmt.Errorf("group not found")
	}

	return result.Error
}

func (gr *groupRepository) IsUserAdmin(userId, groupId uint) (bool, error) {
	var count int64

	err := gr.db.
		Table("group_member").
		Where("group_id = ? AND user_id = ? AND invite_status_id = ? AND role_id IN ?",
			groupId, userId, acceptedStatusId, []uint{ownerId, adminId}).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
