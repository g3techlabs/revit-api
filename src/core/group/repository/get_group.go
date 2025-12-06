package repository

import (
	"github.com/g3techlabs/revit-api/src/core/group/response"
	"gorm.io/gorm"
)

func (gr *groupRepository) GetGroup(userId, groupId uint) (*response.GroupResponse, error) {
	var group response.GroupResponse

	query := gr.buildGetGroupBaseQuery(userId, acceptedStatusId)
	query = query.Where("g.id = ?", groupId)

	if err := query.Scan(&group).Error; err != nil {
		return nil, err
	}

	return &group, nil
}

func (gr *groupRepository) buildGetGroupBaseQuery(userId uint, acceptedStatusId uint) *gorm.DB {

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
					jsonb_build_object(
						'nickname', u.nickname,
						'profilePicUrl', CAST(? AS text) || u.profile_pic
					)
				) FILTER (WHERE u.user_id IS NOT NULL),
				'[]'
			) AS friends_in_group,
			COALESCE(
				json_agg(
					jsonb_build_object(
						'name', e.name,
						'date', e.date,
						'coordinates', jsonb_build_object(
							'latitude', ST_Y(e.location::geometry),
							'longitude', ST_X(e.location::geometry)
						),
						'id', e.id,
						'photo', CAST(? AS text) || e.photo
					)
				) FILTER (WHERE e.id IS NOT NULL),
				'[]'
			) AS upcoming_events,
			COUNT(gm_total.user_id) AS total_members 
		`, cloudFrontUrl, cloudFrontUrl, cloudFrontUrl, cloudFrontUrl).
		Joins("JOIN visibility v ON v.id = g.visibility_id").
		Joins("JOIN city c ON c.id = g.city_id").
		Joins("JOIN state s ON s.id = c.state_id").
		Joins("LEFT JOIN group_member gm_user ON gm_user.group_id = g.id AND gm_user.user_id = ? AND invite_status_id = ? AND left_at IS NULL AND removed_by IS NULL", userId, acceptedStatusId).
		Joins("LEFT JOIN role r ON r.id = gm_user.role_id").
		Joins("LEFT JOIN group_member gm_total ON gm_total.group_id = g.id AND gm_total.invite_status_id = ? AND gm_total.left_at IS NULL and gm_total.removed_by IS NULL", acceptedStatusId).
		Joins("LEFT JOIN event e ON e.group_id = g.id AND e.date >= NOW() AND e.canceled = FALSE").
		Joins("LEFT JOIN LATERAL (?) AS u ON TRUE", friendsSubquery).
		Group("g.id, v.name, c.name, s.name, r.name").
		Where("(g.visibility_id = ? OR gm_user.user_id IS NOT NULL)", PublicVisibility)
}
