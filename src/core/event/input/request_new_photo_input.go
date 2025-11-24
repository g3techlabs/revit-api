package input

// RequestNewPhotoInput representa os dados para solicitar upload de nova foto
// @Description Dados para solicitar URL pré-assinada para upload de foto do evento
type RequestNewPhotoInput struct {
	// Tipo de conteúdo da foto (image/jpeg, image/png, image/webp)
	PhotoContentType string `json:"photoContentType" validate:"required,oneof=image/jpeg image/png image/webp" example:"image/jpeg"`
}
