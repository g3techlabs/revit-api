package repository

import (
	"math"

	"github.com/g3techlabs/revit-api/src/core/group/input"
	"github.com/g3techlabs/revit-api/src/core/group/response"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (gr *groupRepository) GetAdminGroups(userId uint, queryParams input.GetAdminGroupsInput) (*response.GetAdminGroupsResponse, error) {
	limit := int(queryParams.Limit)
	page := int(queryParams.Page)
	offset := (page - 1) * limit

	countQuery := gr.db.Table("groups g").
		Select("g.id").
		Joins("INNER JOIN group_member gm ON gm.group_id = g.id").
		Where("gm.user_id = ? AND gm.invite_status_id = ? AND gm.role_id IN ? AND gm.left_at IS NULL AND gm.removed_by IS NULL",
			userId, acceptedStatusId, []uint{ownerRoleId, adminRoleId})

	if queryParams.Name != "" {
		countQuery = countQuery.Where("LOWER(g.name) LIKE LOWER(?)", "%"+queryParams.Name+"%")
	}

	var totalCount int64
	if err := gr.db.Raw("SELECT COUNT(*) FROM (?) AS subquery", countQuery).Scan(&totalCount).Error; err != nil {
		return nil, err
	}

	totalPages := uint(0)
	if totalCount > 0 && limit > 0 {
		totalPages = uint(math.Ceil(float64(totalCount) / float64(limit)))
	}

	// Query principal
	var groups []response.AdminGroup
	mainQuery := gr.db.Table("groups g").
		Select("g.id, g.name").
		Joins("INNER JOIN group_member gm ON gm.group_id = g.id").
		Where("gm.user_id = ? AND gm.invite_status_id = ? AND gm.role_id IN ? AND gm.left_at IS NULL AND gm.removed_by IS NULL",
			userId, acceptedStatusId, []uint{ownerRoleId, adminRoleId})

	if queryParams.Name != "" {
		mainQuery = mainQuery.Where("LOWER(g.name) LIKE LOWER(?)", "%"+queryParams.Name+"%")
	}

	mainQuery = gr.buildOrderByAdminGroups(mainQuery, queryParams.Name)
	mainQuery = mainQuery.Limit(limit).Offset(offset)

	if err := mainQuery.Scan(&groups).Error; err != nil {
		return nil, err
	}

	if groups == nil {
		groups = make([]response.AdminGroup, 0)
	}

	return &response.GetAdminGroupsResponse{
		Groups:      groups,
		CurrentPage: uint(page),
		TotalPages:  totalPages,
	}, nil
}

func (gr *groupRepository) buildOrderByAdminGroups(query *gorm.DB, nameFilter string) *gorm.DB {
	if nameFilter != "" {
		return query.Order(clause.Expr{SQL: "similarity(g.name, ?) DESC, g.name ASC", Vars: []any{nameFilter}})
	}

	return query.Order("g.name ASC")
}
