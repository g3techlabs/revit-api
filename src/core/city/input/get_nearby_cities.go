package input

// GetNearbyCitiesFilters representa os parâmetros de query para buscar cidades próximas
// @Description Parâmetros para buscar cidades próximas a uma localização geográfica
type GetNearbyCitiesFilters struct {
	// Latitude da localização (deve estar entre -90 e 90)
	Latitude float64 `validate:"required,gte=-90,lte=90" example:"-23.5505"`
	// Longitude da localização (deve estar entre -180 e 180)
	Longitude float64 `validate:"required,gte=-180,lte=180" example:"-46.6333"`
}
