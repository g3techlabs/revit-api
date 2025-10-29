package service

import (
	"github.com/g3techlabs/revit-api/core/event/errors"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (es *EventService) InviteUserToEvent(eventAdminId, eventId, invitedId uint) error {
	if eventAdminId == invitedId {
		return errors.UsersAreTheSame()
	}

	if err := es.eventRepo.MakeEventInvitation(eventAdminId, eventId, invitedId); err != nil {
		switch err.Error() {
		case "requester not a event admin":
			return errors.UserNotAdmin()
		case "invite target already invited/subscribed":
			return errors.UserIsAlreadySubscribed()
		default:
			return generics.InternalError()
		}
	}

	return nil
}
