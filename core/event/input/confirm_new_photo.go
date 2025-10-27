package input

type ConfirmNewPhoto struct {
	Key string `json:"key" validate:"required"`
}
