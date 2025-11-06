package service

import (
	"github.com/g3techlabs/revit-api/src/core/geolocation/service"
	"github.com/g3techlabs/revit-api/src/core/route/input"
	"github.com/g3techlabs/revit-api/src/core/route/repository"
	"github.com/g3techlabs/revit-api/src/core/route/response"
	"github.com/g3techlabs/revit-api/src/infra/websocket"
	"github.com/g3techlabs/revit-api/src/validation"
)

type IRouteService interface {
	CreateRoute(userId uint, data *input.CreateRouteInput) (*response.RouteCreatedReponse, error)
	GetOnlineFriendsToInvite(userId uint) (*[]response.OnlineFriendsResponse, error)
	GetNearbyUsersToInvite(userId uint, data *input.GetNearbyUsersToInviteQuery) (*[]response.NearbyUserToRouteResponse, error)
	InviteNearbyUsers(userId, routeId uint, inviteds *input.UsersToInviteInput) error
}

type RouteService struct {
	validator          validation.IValidator
	geoLocationService service.IGeoLocationService
	routeRepo          repository.RouteRepository
	hub                *websocket.Hub
}

func NewRouteService(
	validator validation.IValidator,
	geoLocationService service.IGeoLocationService,
	routeRepo repository.RouteRepository,
	hub *websocket.Hub,
) IRouteService {
	return &RouteService{
		validator:          validator,
		routeRepo:          routeRepo,
		geoLocationService: geoLocationService,
		hub:                hub,
	}
}
