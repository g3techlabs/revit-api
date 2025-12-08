package repository

import (
	"fmt"

	"github.com/g3techlabs/revit-api/src/db/models"
	"gorm.io/gorm"
)

func (er *eventRepository) CancelEvent(userId, eventId uint) error {
	var event models.Event
	if err := er.db.Select("id").Where("id = ? AND canceled = FALSE", eventId).First(&event).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("event not found or already canceled")
		}
		return err
	}

	var count int64
	err := er.db.
		Table("event_subscriber").
		Where("event_id = ? AND user_id = ? AND invite_status_id = ? AND role_id = ?",
			eventId, userId, acceptedStatusId, ownerRoleId).
		Where("left_at IS NULL AND removed_by IS NULL").
		Count(&count).Error

	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("user is not the event owner")
	}

	result := er.db.Model(&models.Event{}).
		Where("id = ?", eventId).
		Update("canceled", true)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("event not found")
	}

	return nil
}
