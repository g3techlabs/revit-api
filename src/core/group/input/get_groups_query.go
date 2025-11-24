package input

// GetGroupsQuery representa os parâmetros de query para buscar grupos
// @Description Parâmetros de filtro e paginação para listar grupos
type GetGroupsQuery struct {
	// Nome do grupo para busca (opcional)
	Name string `validate:"omitempty" example:"Ciclistas"`
	// ID da cidade para filtrar (opcional)
	CityId uint `validate:"omitempty,number,gt=0" example:"1"`
	// ID do estado para filtrar (opcional)
	StateId uint `validate:"omitempty,number,gt=0" example:"1"`
	// Visibilidade do grupo (opcional: public ou private)
	Visibility string `validate:"omitempty,oneof=public private" example:"public"`
	// Filtrar apenas grupos onde o usuário é membro (opcional)
	Member bool `example:"true"`
	// Limite de resultados por página (opcional, máximo 20)
	Limit uint `validate:"omitempty,number,gt=0,max=20" example:"10"`
	// Número da página (opcional, deve ser maior que 0)
	Page uint `validate:"omitempty,number,gt=0" example:"1"`
}
