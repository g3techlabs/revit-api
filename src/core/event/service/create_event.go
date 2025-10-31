package service

import (
	"fmt"
	"time"

	"github.com/g3techlabs/revit-api/src/core/event/errors"
	"github.com/g3techlabs/revit-api/src/core/event/input"
	"github.com/g3techlabs/revit-api/src/core/event/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
	"github.com/g3techlabs/revit-api/src/utils"
)

func (es *EventService) CreateEvent(userId uint, input *input.CreateEventInput) (*response.PresginedEventPhotoResponse, error) {
	if err := es.validator.Validate(input); err != nil {
		return nil, err
	}

	if input.GroupID != nil {
		isUserGroupAdmin, err := es.eventRepo.IsUserGroupAdmin(userId, *input.GroupID)
		if err != nil {
			return nil, generics.InternalError()
		}
		if !isUserGroupAdmin {
			return nil, generics.Forbidden("User does not have permission of managing this group")
		}
	}

	_, err := es.validateEventDate(input.Date)
	if err != nil {
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

func (es *EventService) validateEventDate(dateString string) (*time.Time, error) {
	formattedDate, err := time.Parse("2006-01-02T15:04:05Z07:00", dateString)
	if err != nil {
		return nil, errors.InvalidDateFormat()
	}

	fifteenMinutesFromNow := time.Now().Add(time.Minute * 15)
	if formattedDate.Before(fifteenMinutesFromNow) {
		return nil, errors.InvalidDateValue()
	}

	return &formattedDate, nil
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
		return generics.InvalidFileExtension()
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
