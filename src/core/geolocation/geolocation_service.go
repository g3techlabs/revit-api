package geolocation

import (
	"encoding/json"

	"github.com/g3techlabs/revit-api/src/core/geolocation/input"
	"github.com/g3techlabs/revit-api/src/core/geolocation/repository"
	"github.com/g3techlabs/revit-api/src/infra/websocket"
	"github.com/g3techlabs/revit-api/src/infra/websocket/response"
	"github.com/g3techlabs/revit-api/src/validation"
)

type IGeoLocationService interface {
	PutUserLocation(userId uint, data *input.Coordinates) error
}

type GeoLocationService struct {
	geoLocationRepository repository.IGeoLocationRepository
	hub                   *websocket.Hub
	validator             validation.IValidator
}

func NewGeoLocationService(validator validation.IValidator, repository repository.IGeoLocationRepository, hub *websocket.Hub) IGeoLocationService {
	return &GeoLocationService{
		validator:             validator,
		geoLocationRepository: repository,
		hub:                   hub,
	}
}

func (gls *GeoLocationService) PutUserLocation(userId uint, data *input.Coordinates) error {
	if err := gls.validator.Validate(data); err != nil {
		return err
	}

	targetIDs, err := gls.geoLocationRepository.PutUserLocation(userId, data)
	if err != nil {
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
		return err
	}

	multicastMessage := &websocket.MulticastMessage{
		Payload:       payloadBytes,
		TargetUserIDs: targetIDs,
	}

	gls.hub.Multicast <- multicastMessage

	return nil
}
