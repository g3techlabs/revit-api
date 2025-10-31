package service

import (
	"github.com/g3techlabs/revit-api/src/core/event/errors"
	"github.com/g3techlabs/revit-api/src/core/event/input"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (es *EventService) AnswerPendingInvite(userId, eventId uint, answer *input.PendingInviteAnswer) error {
	if err := es.validator.Validate(answer); err != nil {
		return err
	}

	switch answer.Answer {
	case "accept":
		if err := es.eventRepo.AcceptPendingInvite(eventId, userId); err != nil {
			if err.Error() == "event invite not found" {
				return errors.EventInviteNotFound()
			}
			return generics.InternalError()
		}
	case "reject":
		if err := es.eventRepo.RejectPendingInvite(eventId, userId); err != nil {
			if err.Error() == "event invite not found" {
				return errors.EventInviteNotFound()
			}
			return generics.InternalError()
		}
	}

	return nil
}
