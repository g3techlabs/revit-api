package repository

import (
	"fmt"
	"time"

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
	IsUserAdmin(userId, eventId uint) (bool, error)
	IsUserGroupAdmin(userId, groupId uint) (bool, error)
	UpdateEvent(userId, eventId uint, newDate *time.Time, data *input.UpdateEventInput) error
	InsertNewEventSubscriber(userId, eventId uint) error
	RevokeEventSubscription(userId, eventId uint) error
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

const publicVisibility uint = 1
const privateVisibility uint = 2

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

func (er *eventRepository) IsUserAdmin(userId, eventId uint) (bool, error) {
	var count int64

	err := er.db.
		Table("event_subscriber").
		Where("event_id = ? AND user_id = ? AND invite_status_id = ? AND role_id IN ?",
			eventId, userId, acceptedStatusId, []uint{ownerRoleId, adminRoleId}).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (er *eventRepository) UpdateEvent(userId, eventId uint, newDate *time.Time, data *input.UpdateEventInput) error {
	updates := make(map[string]any)

	if data.Name != nil {
		updates["name"] = *data.Name
	}
	if data.Description != nil {
		updates["description"] = *data.Description
	}
	if newDate != nil {
		updates["date"] = *newDate
	}
	if data.Visibility != nil {
		switch *data.Visibility {
		case "public":
			updates["visibility_id"] = publicVisibility
		case "private":
			updates["visibility_id"] = privateVisibility
		}
	}

	if data.GroupID != nil {
		updates["group_id"] = *data.GroupID
	}

	if data.Location != nil && data.CityID != nil {
		updates["location"] = gorm.Expr("ST_SetSRID(ST_MakePoint(?, ?), 4326)", data.Location.Longitude, data.Location.Latitude)
		updates["city_id"] = *data.CityID
	}

	query := er.db.Model(&models.Event{}).
		Where("id = ?",
			eventId).
		Where(`EXISTS (SELECT 1 
		FROM event_subscriber 
		WHERE event_subscriber.event_id = ? 
		AND event_subscriber.user_id = ? 
		AND event_subscriber.invite_status_id = ? 
		AND event_subscriber.role_id IN ?)
		`,
			eventId, userId, acceptedStatusId, []uint{ownerRoleId, adminRoleId}).
		Updates(updates)

	if query.RowsAffected == 0 {
		return fmt.Errorf("user not admin")
	}

	return query.Error
}

func (er *eventRepository) IsUserGroupAdmin(userId, groupId uint) (bool, error) {
	var count int64

	err := er.db.
		Table("group_member").
		Where("group_id = ? AND user_id = ? AND invite_status_id = ? AND role_id IN ?",
			groupId, userId, acceptedStatusId, []uint{ownerRoleId, adminRoleId}).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (gr *eventRepository) InsertNewEventSubscriber(userId, eventId uint) error {
	return gr.db.Transaction(func(tx *gorm.DB) error {
		var e models.Event
		if err := tx.Select("id", "visibility_id").Where("id = ?", eventId).First(&e).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return fmt.Errorf("event not found")
			}
			return err
		}

		if e.VisibilityID != publicVisibility {
			return fmt.Errorf("event is private")
		}

		var count int64
		if err := tx.Model(&models.EventSubscriber{}).
			Where("event_id = ? AND user_id = ? AND invite_status_id = ? AND left_at IS NULL AND removed_by IS NULL", eventId, userId, acceptedStatusId).
			Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			return fmt.Errorf("user already subscribed")
		}

		if err := gr.registerSubscriber(eventId, userId, tx); err != nil {
			return err
		}

		return nil
	})
}

func (*eventRepository) registerSubscriber(eventId, userId uint, tx *gorm.DB) error {
	es := models.EventSubscriber{
		EventID:        eventId,
		UserID:         userId,
		RoleID:         memberRoleId,
		InviteStatusID: acceptedStatusId,
	}

	err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "event_id"}, {Name: "user_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"role_id":          memberRoleId,
			"invite_status_id": acceptedStatusId,
			"left_at":          nil,
			"removed_by":       nil,
		}),
	}).Create(&es).Error

	return err
}

func (er *eventRepository) RevokeEventSubscription(userId, eventId uint) error {
	var quitter models.EventSubscriber

	if err := er.db.Model(&models.EventSubscriber{}).Select("role_id").
		Where("event_id = ? AND user_id = ? AND invite_status_id = ?", eventId, userId, acceptedStatusId).
		Where("left_at IS NULL AND removed_by IS NULL").
		First(&quitter).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("user is not subscribed")
		}
		return err
	}

	if quitter.RoleID == ownerRoleId {
		return fmt.Errorf("user is owner")
	}

	result := er.db.
		Model(&models.EventSubscriber{}).
		Where("event_id = ? AND user_id = ? AND invite_status_id = ? AND left_at IS NULL AND removed_by IS NULL", eventId, userId, acceptedStatusId).
		Update("left_at", time.Now().UTC())

	if result.RowsAffected == 0 {
		return fmt.Errorf("user is not subscribed")
	}

	return result.Error
}
