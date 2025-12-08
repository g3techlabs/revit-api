package service

import (
	"github.com/g3techlabs/revit-api/src/core/event/input"
	"github.com/g3techlabs/revit-api/src/core/event/repository"
	"github.com/g3techlabs/revit-api/src/core/event/response"
	"github.com/g3techlabs/revit-api/src/infra/storage"
	"github.com/g3techlabs/revit-api/src/validation"
)

type IEventService interface {
	CreateEvent(userId uint, data *input.CreateEventInput) (*response.PresginedEventPhotoResponse, error)
	ConfirmNewPhoto(userId, eventId uint, data *input.ConfirmNewPhoto) error
	GetEvents(userId uint, filters *input.GetEventsFilters) (*response.GetEventsResponse, error)
	UpdateEvent(userId, eventId uint, data *input.UpdateEventInput) error
	RequestNewPhoto(userId, eventId uint, data *input.RequestNewPhotoInput) (*response.PresginedEventPhotoResponse, error)
	SubscribeToEvent(userId, eventId uint) error
	RevokeEventSubscription(userId, eventId uint) error
	InviteUserToEvent(eventAdminId, eventId, invitedId uint) error
	GetPendingInvites(userId uint, query *input.GetPendingInvitesFilters) (*[]response.GetPendingInvitesResponse, error)
	AnswerPendingInvite(userId, eventId uint, answer *input.PendingInviteAnswer) error
	RemoveSubscriber(eventAdminId, eventId, subscriberId uint) error
	GetEventSubscribers(userId, eventId uint, query *input.GetEventSubscribersInput) (*response.EventSubscribersResponse, error)
	CancelEvent(userId, eventId uint) error
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
