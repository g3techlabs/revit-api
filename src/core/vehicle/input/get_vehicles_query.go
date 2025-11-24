package input

// GetVehiclesParams representa os parâmetros de query para buscar veículos
// @Description Parâmetros de filtro e paginação para listar veículos do usuário
type GetVehiclesParams struct {
	// Limite de resultados por página (opcional, máximo 20)
	Limit uint `validate:"omitempty,number,gt=0,max=20" example:"10"`
	// Número da página (opcional, deve ser maior que 0)
	Page uint `validate:"omitempty,number,gt=0" example:"1"`
	// Apelido do veículo para busca (opcional)
	Nickname string `validate:"omitempty" example:"Minha Moto"`
}
