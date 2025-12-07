package service

import (
	"github.com/g3techlabs/revit-api/src/core/event/input"
	"github.com/g3techlabs/revit-api/src/core/event/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (es *EventService) GetEventSubscribers(userId, eventId uint, query *input.GetEventSubscribersInput) (*response.EventSubscribersResponse, error) {
	if err := es.validator.Validate(query); err != nil {
		return nil, err
	}

	canCheck, err := es.eventRepo.CanUserViewEvent(userId, eventId)
	if err != nil {
		return nil, generics.InternalError()
	}

	if !canCheck {
		return nil, generics.Forbidden("User cannot view event subscribers")
	}

	subscribers, err := es.eventRepo.GetEventSubscribers(eventId, *query)
	if err != nil {
		return nil, generics.InternalError()
	}

	return subscribers, nil
}
