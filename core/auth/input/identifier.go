package input

type Identifier struct {
	Identifier     string `json:"identifier" validate:"required"`
	IdentifierType string `json:"identifierType" validate:"required,oneof=email nickname"`
}
