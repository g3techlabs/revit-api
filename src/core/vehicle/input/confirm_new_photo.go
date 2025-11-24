package input

// ConfirmNewPhoto representa os dados para confirmar upload de nova foto
// @Description Chave da foto que foi enviada com sucesso para confirmar o upload
type ConfirmNewPhoto struct {
	// Chave da foto no storage
	ObjectKey string `json:"objectKey" validate:"required" example:"vehicles/123/main-photo.jpg"`
	// Tipo de foto (main ou feed)
	PhotoType string `json:"photoType" validate:"required,oneof=main feed" example:"main"`
}
