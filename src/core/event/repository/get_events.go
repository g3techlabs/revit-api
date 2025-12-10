package repository

import (
	"math"

	"github.com/g3techlabs/revit-api/src/core/event/input"
	"github.com/g3techlabs/revit-api/src/core/event/response"
	"github.com/g3techlabs/revit-api/src/db/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (er *eventRepository) GetEvents(userId uint, filters *input.GetEventsFilters) (*response.GetEventsResponse, error) {
	limit := 20
	page := 1
	if filters != nil {
		if filters.Limit > 0 {
			limit = int(filters.Limit)
		}
		if filters.Page > 0 {
			page = int(filters.Page)
		}
	}

	countSubquery := er.db.Table("event_subscriber").
		Select("event_id, COUNT(*) AS subscribers_count").
		Where("invite_status_id = ? AND left_at IS NULL AND removed_by IS NULL", acceptedStatusId).
		Group("event_id")

	userRoleSubquery := er.db.Table("event_subscriber es").
		Select("es.event_id, r.name AS role_name").
		Joins("INNER JOIN role r ON r.id = es.role_id").
		Where("es.user_id = ? AND es.invite_status_id = ? AND es.left_at IS NULL AND es.removed_by IS NULL", userId, acceptedStatusId)

	friendsSubscribersSubquery := er.db.Table("event_subscriber es").
		Select("es.event_id, COALESCE(jsonb_agg(jsonb_build_object('id', u.id, 'nickname', u.nickname, 'profilePic', CASE WHEN u.profile_pic IS NULL THEN NULL ELSE '"+cloudFrontUrl+"' || u.profile_pic END)) FILTER (WHERE u.id IS NOT NULL), '[]'::jsonb) AS friends_subscribers").
		Joins("INNER JOIN users u ON u.id = es.user_id").
		Joins("INNER JOIN friendship f ON ((f.requester_id = ? AND f.receiver_id = u.id) OR (f.requester_id = u.id AND f.receiver_id = ?))", userId, userId).
		Where("es.invite_status_id = ? AND es.left_at IS NULL AND es.removed_by IS NULL AND f.invite_status_id = ? AND f.removed_at IS NULL", acceptedStatusId, acceptedStatusId).
		Group("es.event_id")

	hostSubquery := er.db.Table("event_subscriber es_host").
		Select("es_host.event_id, jsonb_build_object('id', u_host.id, 'nickname', u_host.nickname, 'profilePic', CASE WHEN u_host.profile_pic IS NULL THEN NULL ELSE '"+cloudFrontUrl+"' || u_host.profile_pic END) AS host").
		Joins("INNER JOIN users u_host ON u_host.id = es_host.user_id").
		Where("es_host.role_id = ? AND es_host.invite_status_id = ? AND es_host.left_at IS NULL AND es_host.removed_by IS NULL", ownerRoleId, acceptedStatusId)

	countQuery := er.buildGetEventsBaseQuery(countSubquery, userRoleSubquery, friendsSubscribersSubquery, hostSubquery, userId)
	countQuery = er.buildWhereStatement(filters, countQuery)
	countQuery = countQuery.Select("event.id")

	var totalCount int64
	if err := er.db.Raw("SELECT COUNT(*) FROM (?) AS subquery", countQuery).Scan(&totalCount).Error; err != nil {
		return nil, err
	}

	totalPages := uint(0)
	if totalCount > 0 && limit > 0 {
		totalPages = uint(math.Ceil(float64(totalCount) / float64(limit)))
	}

	var events []response.SimpleEvent

	mainQuery := er.buildGetEventsBaseQuery(countSubquery, userRoleSubquery, friendsSubscribersSubquery, hostSubquery, userId)
	mainQuery = er.buildWhereStatement(filters, mainQuery)
	mainQuery = er.buildOrderByStatement(mainQuery, filters)
	mainQuery = er.buildPagination(mainQuery, uint(limit), uint(page))

	if err := mainQuery.Scan(&events).Error; err != nil {
		return nil, err
	}

	if events == nil {
		events = make([]response.SimpleEvent, 0)
	}

	return &response.GetEventsResponse{
		Events:      events,
		CurrentPage: uint(page),
		TotalPages:  totalPages,
	}, nil
}

func (er *eventRepository) buildGetEventsBaseQuery(countSubquery, userRoleSubquery, friendsSubscribersSubquery, hostSubquery *gorm.DB, userId uint) *gorm.DB {
	return er.db.Model(&models.Event{}).
		Select(
			"event.id",
			"event.name",
			"event.description",
			"event.date",
			"v.name AS visibility",
			"COALESCE(es_count.subscribers_count, 0) AS subscribers_count",
			"user_role.role_name AS member_type",
			"jsonb_build_object('latitude', ST_Y(event.location::geometry), 'longitude', ST_X(event.location::geometry)) AS coordinates",
			"jsonb_build_object('id', c.id, 'name', c.name, 'stateName', s.acronym, 'countryName', ct.name) AS address",
			"CASE WHEN event.group_id IS NULL OR g.id IS NULL THEN NULL ELSE jsonb_build_object('name', g.name, 'groupMainPhoto', CASE WHEN g.main_photo IS NULL THEN NULL ELSE '"+cloudFrontUrl+"' || g.main_photo END) END AS \"group\"",
			"CASE WHEN event.photo IS NULL THEN NULL ELSE '"+cloudFrontUrl+"' || event.photo END AS photo",
			"user_role.role_name AS member_role",
			"COALESCE(friends_subscribers.friends_subscribers, '[]'::jsonb) AS friends_subscribers",
			"CASE WHEN event.group_id IS NULL THEN host_data.host ELSE NULL END AS host",
		).
		Joins("LEFT JOIN groups g ON event.group_id = g.id").
		Joins("INNER JOIN visibility v ON v.id = event.visibility_id").
		Joins("INNER JOIN city c ON c.id = event.city_id").
		Joins("INNER JOIN state s ON s.id = c.state_id").
		Joins("INNER JOIN country ct ON ct.id = s.country_id").
		Joins("LEFT JOIN event_subscriber es ON es.event_id = event.id AND es.user_id = ? AND es.invite_status_id = ? AND es.left_at IS NULL AND es.removed_by IS NULL", userId, acceptedStatusId).
		Joins("LEFT JOIN (?) AS es_count ON es_count.event_id = event.id", countSubquery).
		Joins("LEFT JOIN (?) AS user_role ON user_role.event_id = event.id", userRoleSubquery).
		Joins("LEFT JOIN (?) AS friends_subscribers ON friends_subscribers.event_id = event.id", friendsSubscribersSubquery).
		Joins("LEFT JOIN (?) AS host_data ON host_data.event_id = event.id", hostSubquery).
		Where("event.canceled = FALSE AND event.date >= NOW()")
}

func (er *eventRepository) buildWhereStatement(f *input.GetEventsFilters, query *gorm.DB) *gorm.DB {
	if f == nil {
		return query
	}

	if f.Type != "" {
		query = query.Where("user_role.role_name = ?", f.Type)
	}

	if f.Name != "" {
		p := "%" + f.Name + "%"
		query = query.Where("event.name ILIKE ?", p)
	}

	if f.FromDate != "" {
		query = query.Where("event.date >= ?", f.FromDate)
	}
	if f.ToDate != "" {
		query = query.Where("event.date <= ?", f.ToDate)
	}

	if f.CityID != 0 {
		query = query.Where("event.city_id = ?", f.CityID)
	}

	if f.Latitude != nil && f.Longitude != nil {
		dwithinExpr := "ST_DWithin(event.location::geography, ST_SetSRID(ST_MakePoint(?, ?), 4326)::geography, ?)"
		query = query.Where(
			dwithinExpr,
			f.Longitude,
			f.Latitude,
			float64(defaultEventRadiusMeters),
		)
	}

	query = query.Where("(event.visibility_id = ? OR es.user_id IS NOT NULL)", publicVisibility)

	return query
}

func (er *eventRepository) buildOrderByStatement(query *gorm.DB, f *input.GetEventsFilters) *gorm.DB {
	if f != nil && f.Name != "" {
		query = query.Order(clause.Expr{SQL: "similarity(event.name, ?) DESC", Vars: []interface{}{f.Name}})
	} else {
		query = query.Order("event.date ASC")
	}

	return query
}

func (er *eventRepository) buildPagination(query *gorm.DB, limitFilter, pageFilter uint) *gorm.DB {
	if limitFilter > 0 {
		limit := int(limitFilter)
		query = query.Limit(limit)
		if pageFilter > 0 {
			offset := (int(pageFilter) - 1) * limit
			query = query.Offset(offset)
		}
	}

	return query
}
