package service

import (
	"github.com/g3techlabs/revit-api/core/event/input"
	"github.com/g3techlabs/revit-api/core/event/repository"
	"github.com/g3techlabs/revit-api/core/event/response"
	"github.com/g3techlabs/revit-api/core/storage"
	"github.com/g3techlabs/revit-api/validation"
)

type IEventService interface {
	CreateEvent(userId uint, data *input.CreateEventInput) (*response.PresginedEventPhotoResponse, error)
	ConfirmNewPhoto(userId, eventId uint, data *input.ConfirmNewPhoto) error
}

type EventService struct {
	validator      validation.IValidator
	eventRepo      repository.EventRepository
	storageService storage.StorageService
}

func NewEventService(validator validation.IValidator, eventRepo repository.EventRepository, storageService storage.StorageService) IEventService {
	return &EventService{
		validator:      validator,
		eventRepo:      eventRepo,
		storageService: storageService,
	}
}
