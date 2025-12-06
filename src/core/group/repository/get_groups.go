package repository

import (
	"math"

	"github.com/g3techlabs/revit-api/src/core/group/input"
	"github.com/g3techlabs/revit-api/src/core/group/response"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (gr *groupRepository) GetGroups(userId uint, filters *input.GetGroupsQuery) (*response.GetGroupsResponse, error) {
	limit := 20
	page := 1
	if filters.Limit > 0 {
		limit = int(filters.Limit)
	}
	if filters.Page > 0 {
		page = int(filters.Page)
	}

	countQuery := gr.db.Table("groups g").
		Select("g.id").
		Joins("JOIN visibility v ON v.id = g.visibility_id").
		Joins("JOIN city c ON c.id = g.city_id").
		Joins("JOIN state s ON s.id = c.state_id").
		Joins("LEFT JOIN group_member gm_user ON gm_user.group_id = g.id AND gm_user.user_id = ? AND invite_status_id = ?", userId, acceptedStatusId)
	countQuery = gr.buildGetGroupsWhereStatement2(countQuery, filters)

	var totalCount int64
	if err := gr.db.Raw("SELECT COUNT(*) FROM (?) AS subquery", countQuery).Scan(&totalCount).Error; err != nil {
		return nil, err
	}

	totalPages := uint(0)
	if totalCount > 0 && limit > 0 {
		totalPages = uint(math.Ceil(float64(totalCount) / float64(limit)))
	}

	var groups []response.SimpleGroup

	mainQuery := gr.buildGetGroupsBaseQuery(userId, acceptedStatusId)
	mainQuery = gr.buildGetGroupsWhereStatement2(mainQuery, filters)
	mainQuery = gr.buildPagination(mainQuery, limit, page)
	mainQuery = gr.buildOrderBy(mainQuery, filters.Name)

	if err := mainQuery.Scan(&groups).Error; err != nil {
		return nil, err
	}

	if groups == nil {
		groups = make([]response.SimpleGroup, 0)
	}

	return &response.GetGroupsResponse{
		Groups:      groups,
		CurrentPage: uint(page),
		TotalPages:  totalPages,
	}, nil
}

func (gr *groupRepository) buildGetGroupsBaseQuery(userId uint, acceptedStatusId uint) *gorm.DB {

	return gr.db.Table("groups g").
		Select(`
			g.id,
			g.name,
			CAST(? AS text) || g.main_photo AS main_photo,
			CAST(? AS text) || g.banner AS banner,
			r.name AS member_type,
			COUNT(DISTINCT gm_total.user_id) AS total_members
		`, cloudFrontUrl, cloudFrontUrl).
		Joins("JOIN visibility v ON v.id = g.visibility_id").
		Joins("JOIN city c ON c.id = g.city_id").
		Joins("JOIN state s ON s.id = c.state_id").
		Joins("LEFT JOIN group_member gm_user ON gm_user.group_id = g.id AND gm_user.user_id = ? AND invite_status_id = ?", userId, acceptedStatusId).
		Joins("LEFT JOIN role r ON r.id = gm_user.role_id").
		Joins("LEFT JOIN group_member gm_total ON gm_total.group_id = g.id AND gm_total.invite_status_id = ? AND gm_total.left_at IS NULL AND gm_total.removed_by IS NULL", acceptedStatusId).
		Group("g.id, v.name, c.name, s.name, r.name")
}

func (gr *groupRepository) buildGetGroupsWhereStatement2(query *gorm.DB, filters *input.GetGroupsQuery) *gorm.DB {
	if filters == nil {
		return query
	}

	if filters.Name != "" {
		query = query.Where("LOWER(g.name) LIKE LOWER(?)", "%"+filters.Name+"%")
	}
	if filters.CityId != 0 {
		query = query.Where("g.city_id = ?", filters.CityId)
	}
	if filters.StateId != 0 {
		query = query.Where("c.state_id = ?", filters.StateId)
	}
	if filters.Member {
		query = query.Where("gm_user.user_id IS NOT NULL")
	}

	query = query.Where("EXISTS (SELECT 1 FROM group_member WHERE group_member.group_id = g.id AND invite_status_id = ? AND left_at IS NULL AND removed_by IS NULL)", acceptedStatusId)
	query = query.Where("(g.visibility_id = ? OR gm_user.user_id IS NOT NULL)", PublicVisibility)

	return query
}

func (gr *groupRepository) buildPagination(query *gorm.DB, limit, page int) *gorm.DB {

	offset := (page - 1) * limit

	return query.Limit(limit).Offset(offset)
}

func (gr *groupRepository) buildOrderBy(query *gorm.DB, name string) *gorm.DB {
	if name != "" {
		return query.Order(clause.Expr{SQL: "similarity(g.name, ?) DESC", Vars: []any{name}})
	}

	return query.Order("g.created_at DESC")
}
