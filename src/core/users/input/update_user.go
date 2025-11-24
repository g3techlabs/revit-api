package input

// UpdateUser representa os dados para atualização de um usuário
// @Description Dados opcionais para atualizar informações do usuário autenticado
type UpdateUser struct {
	// Nome completo do usuário (opcional)
	Name string `json:"name" validate:"omitempty" example:"João Silva"`
	// Data de nascimento no formato YYYY-MM-DD (opcional)
	Birthdate *string `json:"birthdate" validate:"omitempty,datetime=2006-01-02" example:"1990-01-15"`
	// Flag para remover a foto de perfil (opcional)
	RemoveProfilePic *bool `json:"removeProfilePic" validate:"omitempty" example:"false"`
}
