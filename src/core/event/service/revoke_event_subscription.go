package service

import (
	"github.com/g3techlabs/revit-api/src/core/event/errors"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (es *EventService) RevokeEventSubscription(userId, eventId uint) error {
	if err := es.eventRepo.RevokeEventSubscription(userId, eventId); err != nil {
		switch err.Error() {
		case "user is not subscribed":
			return errors.UserIsNotSubscribed()
		case "user is owner":
			return errors.UserIsEventOwner()
		default:
			return generics.InternalError()
		}
	}

	return nil
}
