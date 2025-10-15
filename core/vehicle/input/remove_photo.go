package input

type RemovePhoto struct {
	PhotoType string `json:"photoType" validate:"required,oneof=main feed"`
}
