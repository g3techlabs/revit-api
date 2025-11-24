package input

// LoginCredentials representa as credenciais de login do usuário
// @Description Credenciais para autenticação do usuário
type LoginCredentials struct {
	// Identificador do usuário (email ou nickname)
	Identifier string `json:"identifier" validate:"required" example:"usuario@email.com"`
	// Senha do usuário
	Password string `json:"password" validate:"required" example:"senha123"`
	// Tipo de identificador usado (email ou nickname)
	IdentifierType string `json:"identifierType" validate:"required,oneof=email nickname" example:"email"`
}
