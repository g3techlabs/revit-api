package input

// RequestPhotoUpsert representa os dados para solicitar upload de foto
// @Description Dados para solicitar URL pré-assinada para upload de foto principal ou de feed do veículo
type RequestPhotoUpsert struct {
	// Tipo de conteúdo da foto (image/jpeg, image/png, image/webp)
	ContentType string `json:"contentType" validate:"required,oneof=image/jpeg image/png image/webp" example:"image/jpeg"`
	// Tipo de foto (main ou feed)
	PhotoType string `json:"photoType" validate:"required,oneof=main feed" example:"main"`
}
