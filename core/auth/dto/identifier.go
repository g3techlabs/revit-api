package dto

type Identifier struct {
	Identifier string `json:"identifier" validate:"required"`
}
