package service

import (
	"github.com/g3techlabs/revit-api/core/event/input"
	"github.com/g3techlabs/revit-api/core/event/response"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (es *EventService) RequestNewPhoto(userId, eventId uint, data *input.RequestNewPhotoInput) (*response.PresginedEventPhotoResponse, error) {
	if err := es.validator.Validate(data); err != nil {
		return nil, err
	}

	isUserAdmin, err := es.eventRepo.IsUserAdmin(userId, eventId)
	if err != nil {
		return nil, generics.InternalError()
	}

	if !isUserAdmin {
		return nil, generics.Forbidden("User does not have permission of editing this group")
	}

	response, err := es.buildResponse(eventId, &data.PhotoContentType)
	if err != nil {
		return nil, err
	}

	return response, nil
}
