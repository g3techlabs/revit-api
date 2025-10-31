package service

import (
	"github.com/g3techlabs/revit-api/src/core/event/errors"
	"github.com/g3techlabs/revit-api/src/core/event/input"
	"github.com/g3techlabs/revit-api/src/response/generics"
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

	formattedDate, err := es.validateEventDate(*data.Date)
	if err != nil {
		return err
	}

	if err := es.eventRepo.UpdateEvent(userId, eventId, formattedDate, data); err != nil {
		if err.Error() == "user not admin" {
			return errors.UserNotAdmin()
		}
		return generics.InternalError()
	}

	return nil
}
