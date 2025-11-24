package input

// GetCitiesFilters representa os parâmetros de query para buscar cidades
// @Description Parâmetros de filtro e paginação para listar cidades por nome
type GetCitiesFilters struct {
	// Nome da cidade para busca
	Name string `validate:"required" example:"São Paulo"`
	// Número da página (opcional, deve ser maior que 0)
	Page uint `validate:"omitempty,number,gt=0" example:"1"`
	// Limite de resultados por página (opcional, deve ser maior que 0)
	Limit uint `validate:"omitempty,number,gt=0" example:"10"`
}
