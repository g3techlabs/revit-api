package input

// Identifier representa um identificador de usuário (email ou nickname)
// @Description Identificador para localizar um usuário no sistema
type Identifier struct {
	// Identificador do usuário (email ou nickname)
	Identifier string `json:"identifier" validate:"required" example:"usuario@email.com"`
	// Tipo de identificador usado (email ou nickname)
	IdentifierType string `json:"identifierType" validate:"required,oneof=email nickname" example:"email"`
}
