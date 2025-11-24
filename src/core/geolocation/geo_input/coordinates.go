package geoinput

// Coordinates representa coordenadas geográficas (latitude e longitude)
// @Description Coordenadas geográficas para localização de pontos em uma rota
type Coordinates struct {
	// Latitude da localização (deve estar entre -85.05112878 e 85.05112878)
	Lat float64 `json:"lat" validate:"required,number,gte=-85.05112878,lte=85.05112878" example:"-23.5505"`
	// Longitude da localização
	Long float64 `json:"long" validate:"required,longitude" example:"-46.6333"`
}
