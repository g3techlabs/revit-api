package service

import (
	"github.com/g3techlabs/revit-api/src/core/event/errors"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (es *EventService) CancelEvent(userId, eventId uint) error {
	err := es.eventRepo.CancelEvent(userId, eventId)
	if err != nil {
		if err.Error() == "event not found or already canceled" {
			return errors.EventNotFound()
		}
		if err.Error() == "user is not the event owner" {
			return generics.Forbidden("Only the event owner can cancel the event")
		}
		return generics.InternalError()
	}

	return nil
}
