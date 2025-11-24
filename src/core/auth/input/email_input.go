package input

// EmailInput representa o email para verificação de disponibilidade
// @Description Email único do usuário para verificar se está disponível
type EmailInput struct {
	// Email único do usuário
	Email string `json:"email" validate:"required,email" example:"joao@email.com"`
}
