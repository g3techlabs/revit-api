package response

// NearbyUserToRouteResponse representa um usuário próximo que pode ser convidado para uma rota
// @Description Informações de um usuário próximo disponível para convite
type NearbyUserToRouteResponse struct {
	// ID do usuário
	ID uint `json:"id" example:"15"`
	// Apelido do usuário
	Nickname string `json:"nickname" example:"mariasantos"`
	// URL da foto de perfil do usuário
	ProfilePic string `json:"profilePic" example:"https://example.com/profile.jpg"`
}
