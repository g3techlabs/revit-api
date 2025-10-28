package repository

import (
	"fmt"

	"github.com/g3techlabs/revit-api/config"
	"github.com/g3techlabs/revit-api/core/event/input"
	"github.com/g3techlabs/revit-api/core/event/response"
	"github.com/g3techlabs/revit-api/db"
	"github.com/g3techlabs/revit-api/db/models"
	"github.com/g3techlabs/revit-api/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EventRepository interface {
	CreateEvent(userId uint, data *models.Event) error
	UpdatePhoto(userId, eventId uint, photoKey string) error
	GetEvents(userId uint, filters *input.GetEventsFilters) (*[]response.GetEventResponse, error)
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository() EventRepository {
	return &eventRepository{
		db: db.Db,
	}
}

const acceptedStatusId uint = 1

const ownerRoleId uint = 1
const adminRoleId uint = 2
const memberRoleId uint = 3

var cloudFrontUrl = config.Get("AWS_CLOUDFRONT_URL")
var defaultEventRadiusMeters = config.GetIntVariable("DEFAULT_EVENT_RADIUS_METERS")

func (er *eventRepository) CreateEvent(userId uint, data *models.Event) error {
	return er.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Event{}).Create(data).Error; err != nil {
			utils.Log.Infof("%s", err.Error())
			return err
		}

		eventOwner := models.EventSubscriber{
			EventID:        data.ID,
			UserID:         userId,
			InviteStatusID: acceptedStatusId,
			RoleID:         ownerRoleId,
		}
		if err := tx.Model(&models.EventSubscriber{}).Create(&eventOwner).Error; err != nil {
			return err
		}

		return nil
	})
}

func (er *eventRepository) UpdatePhoto(userId, eventId uint, photoKey string) error {
	result := er.db.Model(&models.Event{}).
		Where("id = ? AND EXISTS (SELECT 1 FROM event_subscriber WHERE event_subscriber.event_id = ? AND event_subscriber.user_id = ? AND event_subscriber.invite_status_id = ? AND event_subscriber.role_id IN ?)",
			eventId, eventId, userId, acceptedStatusId, []uint{ownerRoleId, adminRoleId}).
		Update("photo", photoKey)

	if result.RowsAffected == 0 {
		return fmt.Errorf("event not found or user not allowed")
	}

	return result.Error
}

func (er *eventRepository) GetEvents(userId uint, filters *input.GetEventsFilters) (*[]response.GetEventResponse, error) {
	var events []response.GetEventResponse

	countSubquery := er.db.Table("event_subscriber").
		Select("event_id, COUNT(*) AS subscribers_count").
		Where("invite_status_id = ? AND left_at IS NULL AND removed_by IS NULL", acceptedStatusId).
		Group("event_id")

	userRoleSubquery := er.db.Table("event_subscriber es").
		Select("es.event_id, r.name AS role_name").
		Joins("INNER JOIN role r ON r.id = es.role_id").
		Where("es.user_id = ? AND es.invite_status_id = ? AND es.left_at IS NULL AND es.removed_by IS NULL", userId, acceptedStatusId)

	query := er.buildGetEventsBaseQuery(countSubquery, userRoleSubquery)
	query = er.buildWhereStatement(filters, query, userId)
	query = er.buildOrderByStatement(query, filters)
	query = er.buildPagination(query, filters.Limit, filters.Page)

	if err := query.Scan(&events).Error; err != nil {
		return nil, err
	}

	if events == nil {
		empty := make([]response.GetEventResponse, 0)
		return &empty, nil
	}

	return &events, nil
}

func (er *eventRepository) buildGetEventsBaseQuery(countSubquery, userRoleSubquery *gorm.DB) *gorm.DB {
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
			"jsonb_build_object('id', c.id, 'name', c.name, 'stateName', s.name, 'countryName', ct.name) AS address",
			"jsonb_build_object('name', g.name, 'groupMainPhoto', CASE WHEN g.main_photo IS NULL THEN NULL ELSE'"+cloudFrontUrl+"' || g.main_photo END) AS \"group\"",
			"CASE WHEN event.photo IS NULL THEN NULL ELSE '"+cloudFrontUrl+"' || event.photo END AS photo",
			"user_role.role_name AS member_role",
		).
		Joins("LEFT JOIN groups g ON event.group_id = g.id").
		Joins("INNER JOIN visibility v ON v.id = event.visibility_id").
		Joins("INNER JOIN city c ON c.id = event.city_id").
		Joins("INNER JOIN state s ON s.id = c.state_id").
		Joins("INNER JOIN country ct ON ct.id = s.country_id").
		Joins("LEFT JOIN (?) AS es_count ON es_count.event_id = event.id", countSubquery).
		Joins("LEFT JOIN (?) AS user_role ON user_role.event_id = event.id", userRoleSubquery).
		Where("event.canceled = FALSE")
}

func (er *eventRepository) buildWhereStatement(f *input.GetEventsFilters, query *gorm.DB, userId uint) *gorm.DB {
	if f == nil {
		return query
	}

	if f.MemberType != nil {
		query = query.Where("user_role.role_name = ?", *f.MemberType)
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

	if f.Visibility != "" {
		switch f.Visibility {
		case "public":
			query = query.Where("event.visibility_id = ?", 1)
		case "private":
			query = query.Where("event.visibility_id = ?", 2)
		}
	}

	if f.Visibility == "private" {
		query = query.Where(
			"EXISTS (SELECT 1 FROM event_subscriber es WHERE es.event_id = event.id AND es.user_id = ? AND es.invite_status_id = ? AND es.left_at IS NULL AND es.removed_by IS NULL)",
			userId, acceptedStatusId,
		)
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
