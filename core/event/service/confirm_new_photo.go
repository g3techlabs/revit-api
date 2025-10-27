package service

import (
	"github.com/g3techlabs/revit-api/core/event/errors"
	"github.com/g3techlabs/revit-api/core/event/input"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (es *EventService) ConfirmNewPhoto(userId, eventId uint, data *input.ConfirmNewPhoto) error {
	if err := es.validator.Validate(data); err != nil {
		return err
	}

	if err := es.eventRepo.UpdatePhoto(userId, eventId, data.Key); err != nil {
		if err.Error() == "event not found or user not allowed" {
			return errors.UserNotAdmin()
		}
		return generics.InternalError()
	}

	return nil
}
