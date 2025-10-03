package input

type Identifier struct {
	Identifier string `json:"identifier" validate:"required"`
}
