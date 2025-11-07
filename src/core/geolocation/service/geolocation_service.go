package service

import (
	geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"
	"github.com/g3techlabs/revit-api/src/core/geolocation/repository"
	"github.com/g3techlabs/revit-api/src/core/geolocation/response"
	"github.com/g3techlabs/revit-api/src/infra/websocket"
	"github.com/g3techlabs/revit-api/src/utils"
	"github.com/g3techlabs/revit-api/src/validation"
	"github.com/redis/go-redis/v9"
)

type IGeoLocationService interface {
	PutUserLocation(userId uint, data *geoinput.Coordinates) error
	PutUserOnFreeRoam(userId uint, data *geoinput.Coordinates) error
	PutUserOnRoute(routeId uint, userDetails *response.UserDetails, data *geoinput.Coordinates) error
	RemoveUserLocation(userId uint) error
	CheckUsersAreOnline(userIDs []uint) ([]bool, error)
	GetNearbyUsersToRouteInvite(userId uint, lat, long float64, page, pageSize int) ([]uint, error)
}

type GeoLocationService struct {
	geoLocationRepository repository.IGeoLocationRepository
	hub                   *websocket.Hub
	validator             validation.IValidator
	logger                utils.ILogger
}

func NewGeoLocationService(validator validation.IValidator, repository repository.IGeoLocationRepository, hub *websocket.Hub, logger utils.ILogger) IGeoLocationService {
	return &GeoLocationService{
		validator:             validator,
		geoLocationRepository: repository,
		hub:                   hub,
		logger:                logger,
	}
}

const FREE_ROAM_KEY string = repository.FREE_ROAM_KEY

func (gls *GeoLocationService) RemoveUserLocation(userId uint) error {
	key, err := gls.geoLocationRepository.GetUserStateGeoKey(userId)
	if err != nil && err != redis.Nil {
		gls.logger.Errorf("Error getting user state geo key on DELETE operation: %v", err)
		return err
	}

	if err := gls.geoLocationRepository.RemoveUserLocation(key, userId); err != nil {
		gls.logger.Errorf("Error removing user location: %v", err)
		return err
	}

	if err := gls.geoLocationRepository.ClearUserState(userId); err != nil {
		gls.logger.Errorf("Error clearing user state: %v", err)
		return err
	}

	return nil
}
