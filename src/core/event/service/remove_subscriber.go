package service

import (
	"github.com/g3techlabs/revit-api/src/core/event/errors"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (es *EventService) RemoveSubscriber(eventAdminId, eventId, subscriberId uint) error {
	if eventAdminId == subscriberId {
		return errors.UsersAreTheSame()
	}

	if err := es.eventRepo.RemoveSubscriber(eventAdminId, eventId, subscriberId); err != nil {
		switch err.Error() {
		case "subscriber not found":
			return errors.UserIsNotSubscribed()
		case "admin not found":
			return errors.UserNotAdmin()
		}
		return generics.InternalError()
	}

	return nil
}
