package response

// FriendshipRequest representa uma solicitação de amizade pendente
// @Description Informações sobre uma solicitação de amizade pendente
type FriendshipRequest struct {
	// ID do usuário que enviou a solicitação
	RequesterID uint `json:"id" example:"42"`
	// Apelido do usuário que enviou a solicitação
	Nickname string `json:"nickname" example:"joaosilva"`
	// URL da foto de perfil do usuário (opcional)
	ProfilePic *string `json:"profilePicUrl" example:"https://example.com/users/42/profile.jpg"`
}
