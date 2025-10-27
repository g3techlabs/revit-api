package service

import (
	"github.com/g3techlabs/revit-api/core/event/input"
	"github.com/g3techlabs/revit-api/core/event/response"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (es *EventService) CreateEvent(userId uint, input *input.CreateEventInput) (*response.PresginedEventPhotoResponse, error) {
	if err := es.validator.Validate(input); err != nil {
		return nil, err
	}

	if err := es.eventRepo.CreateEvent(userId, input.ToEventModel()); err != nil {
		return nil, generics.InternalError()
	}

	return nil, nil
}
