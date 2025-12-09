package service

import (
	"github.com/g3techlabs/revit-api/src/core/event/input"
	"github.com/g3techlabs/revit-api/src/core/event/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (es *EventService) GetAdminEvents(userId uint, query *input.GetAdminEventsInput) (*response.GetAdminEventsResponse, error) {
	if err := es.validator.Validate(query); err != nil {
		return nil, err
	}

	eventsResponse, err := es.eventRepo.GetAdminEvents(userId, *query)
	if err != nil {
		return nil, generics.InternalError()
	}

	return eventsResponse, nil
}
