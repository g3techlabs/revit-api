package repository

import (
	"math"

	"github.com/g3techlabs/revit-api/src/core/event/input"
	"github.com/g3techlabs/revit-api/src/core/event/response"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (er *eventRepository) GetAdminEvents(userId uint, queryParams input.GetAdminEventsInput) (*response.GetAdminEventsResponse, error) {
	limit := int(queryParams.Limit)
	page := int(queryParams.Page)
	offset := (page - 1) * limit

	countQuery := er.db.Table("event e").
		Select("e.id").
		Joins("INNER JOIN event_subscriber es ON es.event_id = e.id").
		Where("es.user_id = ? AND es.invite_status_id = ? AND es.role_id IN ? AND es.left_at IS NULL AND es.removed_by IS NULL",
			userId, acceptedStatusId, []uint{ownerRoleId, adminRoleId})

	if queryParams.Name != "" {
		countQuery = countQuery.Where("LOWER(e.name) LIKE LOWER(?)", "%"+queryParams.Name+"%")
	}

	var totalCount int64
	if err := er.db.Raw("SELECT COUNT(*) FROM (?) AS subquery", countQuery).Scan(&totalCount).Error; err != nil {
		return nil, err
	}

	totalPages := uint(0)
	if totalCount > 0 && limit > 0 {
		totalPages = uint(math.Ceil(float64(totalCount) / float64(limit)))
	}

	var events []response.AdminEvent
	mainQuery := er.db.Table("event e").
		Select("e.id, e.name, e.date").
		Joins("INNER JOIN event_subscriber es ON es.event_id = e.id").
		Where("es.user_id = ? AND es.invite_status_id = ? AND es.role_id IN ? AND es.left_at IS NULL AND es.removed_by IS NULL",
			userId, acceptedStatusId, []uint{ownerRoleId, adminRoleId})

	if queryParams.Name != "" {
		mainQuery = mainQuery.Where("LOWER(e.name) LIKE LOWER(?)", "%"+queryParams.Name+"%")
	}

	mainQuery = er.buildOrderByAdminEvents(mainQuery, queryParams.Name)
	mainQuery = mainQuery.Limit(limit).Offset(offset)

	if err := mainQuery.Scan(&events).Error; err != nil {
		return nil, err
	}

	if events == nil {
		events = make([]response.AdminEvent, 0)
	}

	return &response.GetAdminEventsResponse{
		Events:      events,
		CurrentPage: uint(page),
		TotalPages:  totalPages,
	}, nil
}

func (er *eventRepository) buildOrderByAdminEvents(query *gorm.DB, nameFilter string) *gorm.DB {
	if nameFilter != "" {
		return query.Order(clause.Expr{SQL: "similarity(e.name, ?) DESC, e.date ASC", Vars: []any{nameFilter}})
	}

	return query.Order("e.date ASC")
}
