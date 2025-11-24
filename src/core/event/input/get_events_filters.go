package input

// GetEventsFilters representa os parâmetros de query para buscar eventos
// @Description Parâmetros de filtro e paginação para listar eventos
type GetEventsFilters struct {
	// Nome do evento para busca (opcional)
	Name string `validate:"omitempty" example:"Pedal"`
	// Data inicial para filtrar eventos (opcional, formato: 2006-01-02)
	FromDate string `validate:"omitempty,datetime=2006-01-02" example:"2024-12-01"`
	// Data final para filtrar eventos (opcional, formato: 2006-01-02)
	ToDate string `validate:"omitempty,datetime=2006-01-02" example:"2024-12-31"`
	// Latitude para filtrar eventos próximos (opcional, requer Longitude)
	Latitude *float64 `validate:"required_with=Longitude,omitempty,latitude" example:"-23.5505"`
	// Longitude para filtrar eventos próximos (opcional, requer Latitude)
	Longitude *float64 `validate:"required_with=Latitude,omitempty,longitude" example:"-46.6333"`
	// ID da cidade para filtrar (opcional)
	CityID uint `validate:"omitempty,number,gt=0" example:"1"`
	// Tipo de membro do usuário no evento (opcional: owner, admin, member)
	MemberType *string `validate:"omitempty,oneof=owner admin member" example:"member"`
	// Visibilidade do evento (opcional: public ou private)
	Visibility string `validate:"omitempty,oneof=public private" example:"public"`
	// Limite de resultados por página (opcional)
	Limit uint `validate:"omitempty,number,gt=0" example:"10"`
	// Número da página (opcional, deve ser maior que 0)
	Page uint `validate:"omitempty,number,gt=0" example:"1"`
}
