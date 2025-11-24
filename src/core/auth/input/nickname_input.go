package input

// NicknameInput representa o apelido para verificação de disponibilidade
// @Description Apelido único do usuário para verificar se está disponível
type NicknameInput struct {
	// Apelido único do usuário (3-32 caracteres, apenas minúsculas)
	Nickname string `json:"nickname" validate:"required,min=3,max=32,lowercase" example:"joaosilva"`
}
