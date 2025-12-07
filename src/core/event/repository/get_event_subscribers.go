package repository

import (
	"math"

	"github.com/g3techlabs/revit-api/src/core/event/input"
	"github.com/g3techlabs/revit-api/src/core/event/response"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (er *eventRepository) GetEventSubscribers(eventId uint, queryParams input.GetEventSubscribersInput) (*response.EventSubscribersResponse, error) {
	limit := int(queryParams.Limit)
	page := int(queryParams.Page)
	offset := (page - 1) * limit

	countQuery := er.db.Table("event_subscriber es").
		Select("es.user_id").
		Where("es.event_id = ? AND es.invite_status_id = ? AND es.left_at IS NULL AND es.removed_by IS NULL", eventId, acceptedStatusId)
	if queryParams.Nickname != nil && *queryParams.Nickname != "" {
		countQuery = countQuery.Joins("JOIN users u ON u.id = es.user_id").
			Where("LOWER(u.nickname) LIKE LOWER(?)", "%"+*queryParams.Nickname+"%")
	}
	var totalCount int64
	if err := er.db.Raw("SELECT COUNT(*) FROM (?) AS subquery", countQuery).Scan(&totalCount).Error; err != nil {
		return nil, err
	}

	totalPages := uint(0)
	if totalCount > 0 && limit > 0 {
		totalPages = uint(math.Ceil(float64(totalCount) / float64(limit)))
	}

	var subscribers []response.EventSubscriber
	mainQuery := er.db.Table("event_subscriber es").
		Select("u.nickname, u.id AS id, CAST(? AS text) || u.profile_pic AS profile_pic_url, r.name AS role", cloudFrontUrl).
		Joins("INNER JOIN users u ON u.id = es.user_id").
		Joins("INNER JOIN role r ON r.id = es.role_id").
		Where("es.event_id = ? AND es.invite_status_id = ? AND es.left_at IS NULL AND es.removed_by IS NULL", eventId, acceptedStatusId)
	if queryParams.Nickname != nil && *queryParams.Nickname != "" {
		mainQuery = mainQuery.Where("LOWER(u.nickname) LIKE LOWER(?)", "%"+*queryParams.Nickname+"%")
	}

	mainQuery = er.buildOrderByGetEventSubscribers(mainQuery, queryParams.Nickname)
	mainQuery = er.buildPaginationGetEventSubscribers(mainQuery, limit, offset)

	if err := mainQuery.Scan(&subscribers).Error; err != nil {
		return nil, err
	}

	if subscribers == nil {
		subscribers = make([]response.EventSubscriber, 0)
	}

	return &response.EventSubscribersResponse{
		Subscribers: subscribers,
		CurrentPage: uint(page),
		TotalPages:  totalPages,
	}, nil
}

func (er *eventRepository) buildOrderByGetEventSubscribers(query *gorm.DB, nicknameFilter *string) *gorm.DB {
	if nicknameFilter != nil && *nicknameFilter != "" {
		return query.Order(clause.Expr{SQL: "similarity(u.nickname, ?) DESC, u.nickname ASC", Vars: []any{*nicknameFilter}})
	}

	return query.Order("u.nickname ASC")
}

func (er *eventRepository) buildPaginationGetEventSubscribers(query *gorm.DB, limit, offset int) *gorm.DB {
	return query.Limit(limit).Offset(offset)
}
