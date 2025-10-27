package repository

import (
	"github.com/g3techlabs/revit-api/db"
	"github.com/g3techlabs/revit-api/db/models"
	"github.com/g3techlabs/revit-api/utils"
	"gorm.io/gorm"
)

type EventRepository interface {
	CreateEvent(userId uint, data *models.Event) error
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
