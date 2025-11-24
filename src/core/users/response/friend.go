package response

import "time"

// Friend representa um amigo do usuário
// @Description Informações sobre um amigo do usuário autenticado
type Friend struct {
	// ID do amigo
	ID uint `json:"id" example:"15"`
	// Nome completo do amigo
	Name string `json:"name" example:"Maria Santos"`
	// Apelido do amigo
	Nickname string `json:"nickname" example:"mariasantos"`
	// URL da foto de perfil do amigo (opcional)
	ProfilePicUrl *string `json:"ProfilePicUrl" example:"https://example.com/users/15/profile.jpg"`
	// Data de criação da conta do amigo
	Since time.Time `json:"since" example:"2024-01-15T10:30:00Z"`
	// Data em que se tornaram amigos (opcional)
	FriendsSince *time.Time `json:"friendsSince" example:"2024-06-20T14:00:00Z"`
}
