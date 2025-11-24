package input

import geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"

// CreateRouteInput representa os dados para criação de uma nova rota
// @Description Dados necessários para criar uma nova rota com localização de início e destino
type CreateRouteInput struct {
	// Localização inicial da rota (coordenadas geográficas)
	StartLocation geoinput.Coordinates `json:"startLocation" validate:"required"`
	// Localização de destino da rota (coordenadas geográficas)
	Destination geoinput.Coordinates `json:"destination" validate:"required"`
}
