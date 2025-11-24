package input

// ConfirmNewPhotos representa os dados para confirmar upload de novas fotos
// @Description Chaves das fotos que foram enviadas com sucesso para confirmar o upload
type ConfirmNewPhotos struct {
	// Chave da foto principal no storage (opcional)
	MainPhotoKey *string `json:"mainPhotoKey" validate:"omitempty" example:"groups/123/main-photo.jpg"`
	// Chave do banner no storage (opcional)
	BannerKey *string `json:"bannerKey" validate:"omitempty" example:"groups/123/banner.jpg"`
}
