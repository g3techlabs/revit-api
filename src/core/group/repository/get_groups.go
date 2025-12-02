package repository

import (
	"math"

	"github.com/g3techlabs/revit-api/src/core/group/input"
	"github.com/g3techlabs/revit-api/src/core/group/response"
	"gorm.io/gorm"
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

	var groups []response.GroupResponse

	mainQuery := gr.buildGetGroupsBaseQuery(userId, acceptedStatusId)
	mainQuery = gr.buildGetGroupsWhereStatement2(mainQuery, filters)
	mainQuery = gr.buildPagination(mainQuery, limit, page)

	if err := mainQuery.Scan(&groups).Error; err != nil {
		return nil, err
	}

	if groups == nil {
		groups = make([]response.GroupResponse, 0)
	}

	return &response.GetGroupsResponse{
		Groups:      groups,
		CurrentPage: uint(page),
		TotalPages:  totalPages,
	}, nil
}

func (gr *groupRepository) buildGetGroupsBaseQuery(userId uint, acceptedStatusId uint) *gorm.DB {

	friendsSubquery := gr.db.Table("group_member gm_friend").
		Select("gm_friend.user_id, users.nickname, users.profile_pic").
		Joins("JOIN users ON users.id = gm_friend.user_id").
		Joins(`JOIN friendship f ON (
            (f.requester_id = ? AND f.receiver_id = users.id)
            OR (f.receiver_id = ? AND f.requester_id = users.id)
        )`, userId, userId).
		Where("gm_friend.group_id = g.id AND gm_friend.invite_status_id = ?", acceptedStatusId).
		Order("users.nickname").
		Limit(3)

	return gr.db.Table("groups g").
		Select(`
			g.id,
			g.name,
			g.description,
			CAST(? AS text) || g.main_photo AS main_photo,
			CAST(? AS text) || g.banner AS banner,
			g.created_at,
			v.name AS visibility,
			c.name AS city,
			s.name AS state,
			r.name AS member_type,
			COALESCE(
				json_agg(
					DISTINCT jsonb_build_object(
						'nickname', u.nickname,
						'profilePicUrl', CAST(? AS text) || u.profile_pic
					)
				) FILTER (WHERE u.user_id IS NOT NULL),
				'[]'
			) AS friends_in_group,
			COUNT(DISTINCT gm_total.user_id) AS total_members 
		`, cloudFrontUrl, cloudFrontUrl, cloudFrontUrl).
		Joins("JOIN visibility v ON v.id = g.visibility_id").
		Joins("JOIN city c ON c.id = g.city_id").
		Joins("JOIN state s ON s.id = c.state_id").
		Joins("LEFT JOIN group_member gm_user ON gm_user.group_id = g.id AND gm_user.user_id = ? AND invite_status_id = ?", userId, acceptedStatusId).
		Joins("LEFT JOIN role r ON r.id = gm_user.role_id").
		Joins("LEFT JOIN group_member gm_total ON gm_total.group_id = g.id AND gm_total.invite_status_id = ?", acceptedStatusId).
		Joins("LEFT JOIN LATERAL (?) AS u ON TRUE", friendsSubquery).
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
	if filters.Visibility != "" {
		query = query.Where("LOWER(v.name) = LOWER(?)", filters.Visibility)
	}
	if filters.Member {
		query = query.Where("gm_user.user_id IS NOT NULL")
	}

	query = query.Where("EXISTS (SELECT 1 FROM group_member WHERE group_member.group_id = g.id AND left_at IS NULL)")
	query = query.Where("(g.visibility_id = ? OR gm_user.user_id IS NOT NULL)", PublicVisibility)

	return query
}

func (gr *groupRepository) buildPagination(query *gorm.DB, limit, page int) *gorm.DB {

	offset := (page - 1) * limit

	return query.Limit(limit).Offset(offset).Order("g.created_at DESC")
}
