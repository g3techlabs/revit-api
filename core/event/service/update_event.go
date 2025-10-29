package service

import (
	"time"

	"github.com/g3techlabs/revit-api/core/event/errors"
	"github.com/g3techlabs/revit-api/core/event/input"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (es *EventService) UpdateEvent(userId, eventId uint, data *input.UpdateEventInput) error {
	if err := es.validator.Validate(data); err != nil {
		return err
	}

	isUserAdmin, err := es.eventRepo.IsUserAdmin(userId, eventId)
	if err != nil {
		return generics.InternalError()
	}

	if !isUserAdmin {
		return generics.Forbidden("User does not have permission to edit this event")
	}

	if data.GroupID != nil {
		isUserGroupAdmin, err := es.eventRepo.IsUserGroupAdmin(userId, *data.GroupID)
		if err != nil {
			return generics.InternalError()
		}
		if !isUserGroupAdmin {
			return generics.Forbidden("User does not have permission of managing this group")
		}
	}

	var newDate time.Time
	if data.Date != nil {
		newDate, err = time.Parse("2006-01-02T15:04:05Z07:00", *data.Date)
		if err != nil {
			return errors.InvalidDate()
		}
	}

	if err := es.eventRepo.UpdateEvent(userId, eventId, &newDate, data); err != nil {
		if err.Error() == "user not admin" {
			return errors.UserNotAdmin()
		}
		return generics.InternalError()
	}

	return nil
}
