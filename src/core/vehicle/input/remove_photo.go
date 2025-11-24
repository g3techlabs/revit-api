package input

// RemovePhoto representa os dados para remover uma foto de veículo
// @Description Dados para remover uma foto principal ou de feed do veículo
type RemovePhoto struct {
	// Tipo de foto a ser removida (main ou feed)
	PhotoType string `json:"photoType" validate:"required,oneof=main feed" example:"main"`
}
