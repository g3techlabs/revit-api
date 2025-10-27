package repository

import (
	"fmt"

	"github.com/g3techlabs/revit-api/db"
	"github.com/g3techlabs/revit-api/db/models"
	"github.com/g3techlabs/revit-api/utils"
	"gorm.io/gorm"
)

type EventRepository interface {
	CreateEvent(userId uint, data *models.Event) error
	UpdatePhoto(userId, eventId uint, photoKey string) error
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

func (gr *eventRepository) UpdatePhoto(userId, eventId uint, photoKey string) error {
	result := gr.db.Model(&models.Event{}).
		Where("id = ? AND EXISTS (SELECT 1 FROM event_subscriber WHERE event_subscriber.event_id = ? AND event_subscriber.user_id = ? AND event_subscriber.invite_status_id = ? AND event_subscriber.role_id IN ?)",
			eventId, eventId, userId, acceptedStatusId, []uint{ownerRoleId, adminRoleId}).
		Update("photo", photoKey)

	if result.RowsAffected == 0 {
		return fmt.Errorf("event not found or user not allowed")
	}

	return result.Error
}
