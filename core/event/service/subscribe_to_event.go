package service

import (
	"github.com/g3techlabs/revit-api/core/event/errors"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (es *EventService) SubscribeToEvent(userId, eventId uint) error {
	err := es.eventRepo.InsertNewEventSubscriber(userId, eventId)
	if err != nil {
		if err.Error() == "event not found" || err.Error() == "event is private" {
			return errors.EventNotFound()
		}
		if err.Error() == "user already subscribed" {
			return errors.UserIsAlreadySubscribed()
		}
		return generics.InternalError()
	}

	return nil
}
