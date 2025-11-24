package input

// ConfirmNewPhoto representa os dados para confirmar upload de nova foto
// @Description Chave da foto que foi enviada com sucesso para confirmar o upload
type ConfirmNewPhoto struct {
	// Chave da foto no storage
	Key string `json:"key" validate:"required" example:"events/123/photo.jpg"`
}
