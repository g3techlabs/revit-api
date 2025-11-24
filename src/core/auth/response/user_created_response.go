package response

import "time"

// UserCreatedResponse representa a resposta de criação de usuário
// @Description Resposta retornada quando um usuário é criado com sucesso
type UserCreatedResponse struct {
	// Nome completo do usuário
	Name string `json:"name"`
	// Email do usuário
	Email string `json:"email"`
	// Apelido único do usuário
	Nickname string `json:"nickname"`
	// URL da foto de perfil (se houver)
	ProfilePic *string `json:"profilePic"`
	// Data de nascimento (se houver)
	Birthdate *time.Time `json:"birthdate"`
	// Data de criação
	CreatedAt time.Time `json:"createdAt"`
	// Data de última atualização
	UpdatedAt time.Time `json:"updatedAt"`
}
