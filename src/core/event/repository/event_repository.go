package repository

import (
	"fmt"
	"math"
	"time"

	"github.com/g3techlabs/revit-api/src/config"
	"github.com/g3techlabs/revit-api/src/core/event/input"
	"github.com/g3techlabs/revit-api/src/core/event/response"
	"github.com/g3techlabs/revit-api/src/db"
	"github.com/g3techlabs/revit-api/src/db/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EventRepository interface {
	CreateEvent(userId uint, data *models.Event) error
	UpdatePhoto(userId, eventId uint, photoKey string) error
	GetEvents(userId uint, filters *input.GetEventsFilters) (*response.GetEventsResponse, error)
	IsUserAdmin(userId, eventId uint) (bool, error)
	IsUserGroupAdmin(userId, groupId uint) (bool, error)
	UpdateEvent(userId, eventId uint, newDate *time.Time, data *input.UpdateEventInput) error
	InsertNewEventSubscriber(userId, eventId uint) error
	RevokeEventSubscription(userId, eventId uint) error
	MakeEventInvitation(eventAdminId, eventId, invitedId uint) error
	GetPendingInvites(userId, limit, page uint) (*response.GetPendingInvitesPaginatedResponse, error)
	AcceptPendingInvite(eventId uint, userId uint) error
	RejectPendingInvite(eventId uint, userId uint) error
	RemoveSubscriber(eventAdminId, eventId, subscriberId uint) error
	GetEventSubscribers(eventId uint, queryParams input.GetEventSubscribersInput) (*response.EventSubscribersResponse, error)
	CanUserViewEvent(userId, eventId uint) (bool, error)
	CancelEvent(userId, eventId uint) error
	GetAdminEvents(userId uint, queryParams input.GetAdminEventsInput) (*response.GetAdminEventsResponse, error)
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
const pendingStatusId uint = 2
const rejectedStatusId uint = 3

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

func (er *eventRepository) MakeEventInvitation(eventAdminId, eventId, invitedId uint) error {
	var admin models.EventSubscriber
	result := er.db.Model(&admin).
		Where("event_id = ? AND user_id = ? AND invite_status_id = ? AND role_id IN ?", eventId, eventAdminId, acceptedStatusId, []uint{ownerRoleId, adminRoleId}).
		Where("left_at IS NULL AND removed_by IS NULL").
		First(&admin)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return fmt.Errorf("requester not a event admin")
		}
		return result.Error
	}

	var count int64
	result = er.db.Model(&models.EventSubscriber{}).
		Where("event_id = ? AND user_id = ? AND invite_status_id IN ? AND left_at IS NULL AND removed_by IS NULL", eventId, invitedId, []uint{pendingStatusId, acceptedStatusId}).
		Count(&count)
	if result.Error != nil {
		return result.Error
	}

	if count > 0 {
		return fmt.Errorf("invite target already invited/subscribed")
	}

	return er.makeInvitation(eventAdminId, eventId, invitedId)
}

func (er *eventRepository) makeInvitation(eventAdminId, eventId, invitedId uint) error {
	data := models.EventSubscriber{
		EventID:        eventId,
		UserID:         invitedId,
		InviterID:      &eventAdminId,
		RoleID:         memberRoleId,
		InviteStatusID: pendingStatusId,
	}

	err := er.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "event_id"}, {Name: "user_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"inviter_id":       eventAdminId,
			"role_id":          memberRoleId,
			"invite_status_id": pendingStatusId,
			"left_at":          nil,
			"removed_by":       nil,
		}),
	}).Create(&data).Error

	return err
}

func (er *eventRepository) GetPendingInvites(userId, limit, page uint) (*response.GetPendingInvitesPaginatedResponse, error) {
	limitInt := 20
	pageInt := 1
	if limit > 0 {
		limitInt = int(limit)
	}
	if page > 0 {
		pageInt = int(page)
	}

	baseQuery := er.db.Model(&models.EventSubscriber{}).
		Select("event_subscriber.event_id").
		Joins("INNER JOIN event e ON e.id = event_subscriber.event_id AND e.canceled = FALSE AND e.date > NOW()").
		Joins("INNER JOIN users u ON u.id = event_subscriber.inviter_id").
		Where("event_subscriber.user_id = ? AND invite_status_id = ? AND left_at IS NULL AND removed_by IS NULL", userId, pendingStatusId)

	// Contar total de registros
	var totalCount int64
	countQuery := er.db.Raw("SELECT COUNT(*) FROM (?) AS subquery", baseQuery)
	if err := countQuery.Scan(&totalCount).Error; err != nil {
		return nil, err
	}

	// Calcular totalPages
	totalPages := uint(0)
	if totalCount > 0 && limitInt > 0 {
		totalPages = uint(math.Ceil(float64(totalCount) / float64(limitInt)))
	}

	// Buscar os dados paginados
	var invites []response.GetPendingInvitesResponse
	query := er.db.Model(&models.EventSubscriber{}).
		Select("e.id AS event_id", "e.name AS event_name", "'"+cloudFrontUrl+"' || e.photo AS event_photo", "u.nickname AS invited_by").
		Joins("INNER JOIN event e ON e.id = event_subscriber.event_id AND e.canceled = FALSE AND e.date > NOW()").
		Joins("INNER JOIN users u ON u.id = event_subscriber.inviter_id").
		Where("event_subscriber.user_id = ? AND invite_status_id = ? AND left_at IS NULL AND removed_by IS NULL", userId, pendingStatusId).
		Order("e.date ASC")

	if limitInt > 0 {
		offset := (pageInt - 1) * limitInt
		query = query.Limit(limitInt).Offset(offset)
	}

	if err := query.Scan(&invites).Error; err != nil {
		return nil, err
	}

	if invites == nil {
		invites = make([]response.GetPendingInvitesResponse, 0)
	}

	return &response.GetPendingInvitesPaginatedResponse{
		Invites:     invites,
		CurrentPage: uint(pageInt),
		TotalPages:  totalPages,
	}, nil
}

func (er *eventRepository) AcceptPendingInvite(eventId uint, userId uint) error {
	result := er.db.
		Model(&models.EventSubscriber{}).
		Joins("INNER JOIN event e ON e.id = event_id AND e.date > NOW()").
		Where("user_id = ? AND event_id = ? AND invite_status_id = ?", userId, eventId, pendingStatusId).
		Update("invite_status_id", acceptedStatusId)

	if result.RowsAffected == 0 {
		return fmt.Errorf("event invite not found")
	}

	return result.Error
}

func (er *eventRepository) RejectPendingInvite(eventId, userId uint) error {
	result := er.db.
		Model(&models.EventSubscriber{}).
		Joins("INNER JOIN event e ON e.id = event_id AND e.date > NOW()").
		Where("user_id = ? AND event_id = ? AND invite_status_id = ?", userId, eventId, pendingStatusId).
		Update("invite_status_id", rejectedStatusId)

	if result.RowsAffected == 0 {
		return fmt.Errorf("event invite not found")
	}

	return result.Error
}

func (er *eventRepository) RemoveSubscriber(eventAdminId, eventId, subscriberId uint) error {
	var eventAdmin models.EventSubscriber
	if err := er.getEventAdmin(eventAdminId, eventId, &eventAdmin); err != nil {
		return err
	}

	data := map[string]any{
		"removed_by": eventAdminId,
		"left_at":    time.Now().UTC(),
	}

	switch eventAdmin.RoleID {
	case adminRoleId:
		return er.removeMemberRoleSubscriber(subscriberId, eventId, data)
	case ownerRoleId:
		return er.removeSubscriber(subscriberId, eventId, data)
	}

	return nil
}

func (er *eventRepository) getEventAdmin(eventAdminId, eventId uint, model *models.EventSubscriber) error {
	if err := er.db.
		Model(&model).
		Where("user_id = ? AND event_id = ? AND invite_status_id = ? AND role_id IN ?", eventAdminId, eventId, acceptedStatusId, []uint{ownerRoleId, adminRoleId}).
		First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("admin not found")
		}
		return err
	}

	return nil
}

func (er *eventRepository) removeMemberRoleSubscriber(subscriberId, eventId uint, data map[string]any) error {
	result := er.db.Model(&models.EventSubscriber{}).
		Where("user_id = ? AND event_id = ? AND invite_status_id = ? AND left_at IS NULL AND removed_by IS NULL", subscriberId, eventId, acceptedStatusId).
		Where("role_id = ?", memberRoleId).
		Updates(data)

	if result.RowsAffected == 0 {
		return fmt.Errorf("subscriber not found")
	}

	return result.Error
}

func (er *eventRepository) removeSubscriber(subscriberId, eventId uint, data map[string]any) error {
	result := er.db.Model(&models.EventSubscriber{}).
		Where("user_id = ? AND event_id = ? AND invite_status_id = ? AND left_at IS NULL AND removed_by IS NULL", subscriberId, eventId, acceptedStatusId).
		Where("role_id IN ?", []uint{adminRoleId, memberRoleId}).
		Updates(data)

	if result.RowsAffected == 0 {
		return fmt.Errorf("subscriber not found")
	}

	return result.Error
}
