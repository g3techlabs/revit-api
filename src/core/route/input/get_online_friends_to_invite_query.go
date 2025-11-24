package input

// GetOnlineFriendsToInviteQuery representa os parâmetros de query para buscar amigos online
// @Description Parâmetros de paginação para listar amigos online que podem ser convidados para uma rota
type GetOnlineFriendsToInviteQuery struct {
	// Limite de resultados por página (deve ser entre 1 e 99)
	Limit uint `validate:"required,number,gt=0,lt=100" example:"10"`
	// Número da página (deve ser maior que 0)
	Page uint `validate:"required,number,gt=0" example:"1"`
}
