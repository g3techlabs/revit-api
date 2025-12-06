package repository

import (
	"math"

	"github.com/g3techlabs/revit-api/src/core/group/input"
	"github.com/g3techlabs/revit-api/src/core/group/response"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (gr *groupRepository) GetMembers(groupId uint, queryParams input.GetMembersInput) (*response.GroupMembersResponse, error) {
	limit := int(queryParams.Limit)
	page := int(queryParams.Page)
	offset := (page - 1) * limit

	countQuery := gr.db.Table("group_member gm").
		Select("gm.user_id").
		Where("gm.group_id = ? AND gm.invite_status_id = ? AND gm.left_at IS NULL AND gm.removed_by IS NULL", groupId, acceptedStatusId)
	if queryParams.Nickname != nil && *queryParams.Nickname != "" {
		countQuery = countQuery.Joins("JOIN users u ON u.id = gm.user_id").
			Where("LOWER(u.nickname) LIKE LOWER(?)", "%"+*queryParams.Nickname+"%")
	}
	var totalCount int64
	if err := gr.db.Raw("SELECT COUNT(*) FROM (?) AS subquery", countQuery).Scan(&totalCount).Error; err != nil {
		return nil, err
	}

	totalPages := uint(0)
	if totalCount > 0 && limit > 0 {
		totalPages = uint(math.Ceil(float64(totalCount) / float64(limit)))
	}

	var members []response.GroupMember
	mainQuery := gr.db.Table("group_member gm").
		Select("u.nickname, u.id AS id, CAST(? AS text) || u.profile_pic AS profile_pic_url, r.name AS role", cloudFrontUrl).
		Joins("INNER JOIN users u ON u.id = gm.user_id").
		Joins("INNER JOIN role r ON r.id = gm.role_id").
		Where("gm.group_id = ? AND gm.invite_status_id = ? AND gm.left_at IS NULL AND gm.removed_by IS NULL", groupId, acceptedStatusId)
	if queryParams.Nickname != nil && *queryParams.Nickname != "" {
		mainQuery = mainQuery.Where("LOWER(u.nickname) LIKE LOWER(?)", "%"+*queryParams.Nickname+"%")
	}

	mainQuery = gr.buildOrderByGetMembers(mainQuery, queryParams.Nickname)
	mainQuery = gr.buildPaginationGetMembers(mainQuery, limit, offset)

	if err := mainQuery.Scan(&members).Error; err != nil {
		return nil, err
	}

	if members == nil {
		members = make([]response.GroupMember, 0)
	}

	return &response.GroupMembersResponse{
		Members:     members,
		CurrentPage: uint(page),
		TotalPages:  totalPages,
	}, nil
}

func (gr *groupRepository) buildOrderByGetMembers(query *gorm.DB, nicknameFilter *string) *gorm.DB {
	if nicknameFilter != nil && *nicknameFilter != "" {
		return query.Order(clause.Expr{SQL: "similarity(u.nickname, ?) DESC, u.nickname ASC", Vars: []any{*nicknameFilter}})
	}

	return query.Order("u.nickname ASC")
}

func (gr *groupRepository) buildPaginationGetMembers(query *gorm.DB, limit, offset int) *gorm.DB {
	return query.Limit(limit).Offset(offset)
}
