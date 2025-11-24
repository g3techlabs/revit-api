package input

// RequestProfilePicUpdate representa os dados para solicitar upload de nova foto de perfil
// @Description Dados para solicitar URL pré-assinada para upload de foto de perfil
type RequestProfilePicUpdate struct {
	// Tipo de conteúdo da foto (image/jpeg, image/png, image/webp)
	ContentType string `json:"contentType" validate:"required,oneof=image/jpeg image/png image/webp" example:"image/jpeg"`
}
