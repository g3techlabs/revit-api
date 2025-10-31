package service

import (
	"github.com/g3techlabs/revit-api/src/core/event/input"
	"github.com/g3techlabs/revit-api/src/core/event/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (es *EventService) GetEvents(userId uint, query *input.GetEventsFilters) (*[]response.GetEventResponse, error) {
	if err := es.validator.Validate(query); err != nil {
		return nil, err
	}

	events, err := es.eventRepo.GetEvents(userId, query)
	if err != nil {
		return nil, generics.InternalError()
	}

	return events, nil
}
