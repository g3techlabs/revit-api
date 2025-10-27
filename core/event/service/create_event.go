package service

import (
	"fmt"

	"github.com/g3techlabs/revit-api/core/event/input"
	"github.com/g3techlabs/revit-api/core/event/response"
	"github.com/g3techlabs/revit-api/core/vehicle/errors"
	"github.com/g3techlabs/revit-api/response/generics"
	"github.com/g3techlabs/revit-api/utils"
)

func (es *EventService) CreateEvent(userId uint, input *input.CreateEventInput) (*response.PresginedEventPhotoResponse, error) {
	if err := es.validator.Validate(input); err != nil {
		return nil, err
	}

	eventModel := input.ToEventModel()
	if err := es.eventRepo.CreateEvent(userId, eventModel); err != nil {
		return nil, generics.InternalError()
	}

	response, err := es.buildResponse(eventModel.ID, input.PhotoContentType)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (es *EventService) buildResponse(eventId uint, photoContentType *string) (*response.PresginedEventPhotoResponse, error) {
	response := new(response.PresginedEventPhotoResponse)

	if photoContentType == nil {
		return response, nil
	}

	response.EventId = &eventId
	if err := es.makePresignedPhotoUrl(eventId, *photoContentType, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (es *EventService) makePresignedPhotoUrl(eventId uint, contentType string, r *response.PresginedEventPhotoResponse) error {
	const PHOTO_KEY = "events/%d/main%s"

	extension := utils.MapImageMIMEToExtension(contentType)
	if extension == "" {
		return errors.InvalidFileExtension()
	}

	photoKey := fmt.Sprintf(PHOTO_KEY, eventId, extension)
	presignedUrl, err := es.storageService.PresignPutObjectURL(photoKey, contentType)
	if err != nil {
		return generics.InternalError()
	}

	r.Url = &presignedUrl
	r.ObjectKey = &photoKey

	return nil
}
