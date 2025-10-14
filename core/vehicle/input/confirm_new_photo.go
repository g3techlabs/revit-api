package input

type ConfirmNewPhoto struct {
	ObjectKey string `json:"objectKey" validate:"required"`
	PhotoType string `json:"photoType" validate:"required,oneof=main feed"`
}
