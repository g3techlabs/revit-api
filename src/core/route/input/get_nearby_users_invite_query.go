package input

// GetNearbyUsersToInviteQuery representa os parâmetros de query para buscar usuários próximos
// @Description Parâmetros para buscar usuários próximos a uma localização geográfica que podem ser convidados para uma rota
type GetNearbyUsersToInviteQuery struct {
	// Latitude da localização de referência (deve estar entre -85.05112878 e 85.05112878)
	Lat float64 `json:"lat" validate:"required,number,gte=-85.05112878,lte=85.05112878" example:"-23.5505"`
	// Longitude da localização de referência
	Long float64 `json:"long" validate:"required,longitude" example:"-46.6333"`
	// Limite de resultados por página (deve ser entre 1 e 99)
	Limit uint `json:"limit" validate:"required,number,gt=0,lt=100" example:"10"`
	// Número da página (deve ser maior que 0)
	Page uint `json:"page" validate:"required,number,gt=0" example:"1"`
}
