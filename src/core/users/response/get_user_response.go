package response

import "time"

// GetUserResponse representa um usuário retornado na listagem ou detalhamento
// @Description Informações de um usuário retornado na busca/listagem
type GetUserResponse struct {
	// ID do usuário
	ID uint `json:"id" example:"42"`
	// Nome completo do usuário
	Name string `json:"name" example:"João Silva"`
	// Apelido único do usuário
	Nickname string `json:"nickname" example:"joaosilva"`
	// URL da foto de perfil do usuário (se houver)
	ProfilePicUrl *string `json:"profilePicUrl" example:"https://example.com/users/42/profile.jpg"`
	// Data de criação da conta
	Since *time.Time `json:"since" example:"2024-01-15T10:30:00Z"`
}
