package input

// GetUsersQuery representa os parâmetros de query para buscar usuários
// @Description Parâmetros de filtro e paginação para listar usuários
type GetUsersQuery struct {
	// Número da página (opcional, deve ser maior que 0)
	Page uint `validate:"omitempty,number,gt=0" example:"1"`
	// Limite de resultados por página (opcional, máximo 20)
	Limit uint `validate:"omitempty,number,gt=0,max=20" example:"10"`
	// Apelido para busca (opcional, mínimo 3 caracteres)
	Nickname string `validate:"omitempty,min=3" example:"joao"`
}
