package geolocation

import (
	"encoding/json"

	"github.com/g3techlabs/revit-api/src/core/geolocation/input"
	"github.com/g3techlabs/revit-api/src/core/geolocation/repository"
	"github.com/g3techlabs/revit-api/src/infra/websocket"
	"github.com/g3techlabs/revit-api/src/infra/websocket/response"
	"github.com/g3techlabs/revit-api/src/utils"
	"github.com/g3techlabs/revit-api/src/validation"
)

type IGeoLocationService interface {
	PutUserLocation(userId uint, data *input.Coordinates) error
	RemoveUserLocation(userId uint) error
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

func (gls *GeoLocationService) PutUserLocation(userId uint, data *input.Coordinates) error {
	gls.logger.Info("PutUserLocation operation started")
	if err := gls.validator.Validate(data); err != nil {
		gls.logger.Errorf("%v", err)
		return err
	}

	targetIDs, err := gls.geoLocationRepository.PutUserLocation(userId, data)
	if err != nil {
		gls.logger.Errorf("%v", err)
		return err
	}

	newPayload := &response.UserMovedEvent{
		Event: "user-moved",
		Payload: response.UserMovedPayload{
			Lat:    data.Lat,
			Lng:    data.Long,
			UserID: userId,
		},
	}

	payloadBytes, err := json.Marshal(newPayload)
	if err != nil {
		gls.logger.Errorf("%v", err)
		return err
	}

	multicastMessage := &websocket.MulticastMessage{
		Payload:       payloadBytes,
		TargetUserIDs: targetIDs,
	}

	gls.hub.Multicast <- multicastMessage

	return nil
}

func (gls *GeoLocationService) RemoveUserLocation(userId uint) error {
	if err := gls.geoLocationRepository.RemoveUserLocation(userId); err != nil {
		gls.logger.Errorf("%v", err)
		return err
	}

	return nil
}
