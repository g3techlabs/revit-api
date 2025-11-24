package input

// GetPendingInvites representa os parâmetros de query para listar convites pendentes
// @Description Parâmetros de paginação para listar convites pendentes de grupos
type GetPendingInvites struct {
	// Número da página (opcional, deve ser maior que 0)
	Page uint `validate:"omitempty,number,gt=0" example:"1"`
	// Limite de resultados por página (opcional, máximo 20)
	Limit uint `validate:"omitempty,number,gt=0,max=20" example:"10"`
}
