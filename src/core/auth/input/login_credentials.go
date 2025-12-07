package input

type LoginCredentials struct {
	Identifier     string `json:"identifier" validate:"required" `
	Password       string `json:"password" validate:"required"`
	IdentifierType string `json:"identifierType" validate:"required,oneof=email nickname"`
}
