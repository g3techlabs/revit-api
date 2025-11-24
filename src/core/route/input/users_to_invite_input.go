package input

// UsersToInviteInput representa a entrada de dados para convidar usuários para uma rota
// @Description Lista de IDs dos usuários que serão convidados para participar da rota
type UsersToInviteInput struct {
	// IDs dos usuários a serem convidados (mínimo 1 usuário, cada ID deve ser maior que 0)
	IdsToInvite []uint `json:"idsToInvite" validate:"required,min=1,dive,number,gt=0"`
}
