package response

// OnlineFriendsResponse representa um amigo online que pode ser convidado para uma rota
// @Description Informações de um amigo online disponível para convite
type OnlineFriendsResponse struct {
	// ID do amigo
	FriendId uint `json:"friendId" example:"42"`
	// Apelido do amigo
	Nickname string `json:"nickname" example:"joaosilva"`
	// URL da foto de perfil do amigo
	ProfilePic string `json:"profilePic" example:"https://example.com/profile.jpg"`
}
