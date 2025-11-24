package input

// RequestNewGroupPhotos representa os dados para solicitar upload de novas fotos
// @Description Dados para solicitar URLs pré-assinadas para upload de foto principal e/ou banner do grupo
type RequestNewGroupPhotos struct {
	// Tipo de conteúdo da foto principal (opcional: image/jpeg, image/png, image/webp)
	MainPhotoContentType *string `json:"mainPhotoContentType" validate:"omitempty,oneof=image/jpeg image/png image/webp" example:"image/jpeg"`
	// Tipo de conteúdo do banner (opcional: image/jpeg, image/png, image/webp)
	BannerContentType *string `json:"bannerContentType" validate:"omitempty,oneof=image/jpeg image/png image/webp" example:"image/png"`
}
