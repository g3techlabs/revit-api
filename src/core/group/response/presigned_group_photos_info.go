package response

// PresignedGroupPhotosInfo representa as URLs pré-assinadas para upload de fotos
// @Description URLs pré-assinadas e chaves para upload de foto principal e banner do grupo
type PresignedGroupPhotosInfo struct {
	// URL pré-assinada para upload da foto principal (opcional)
	PresignedMainPhotoUrl *string `json:"presignedMainPhotoUrl" example:"https://s3.amazonaws.com/bucket/groups/123/main.jpg?X-Amz-Signature=..."`
	// Chave da foto principal no storage (opcional)
	MainPhotoKey *string `json:"mainPhotoKey" example:"groups/123/main-photo.jpg"`
	// URL pré-assinada para upload do banner (opcional)
	PresignedBannerUrl *string `json:"presignedBannerUrl" example:"https://s3.amazonaws.com/bucket/groups/123/banner.jpg?X-Amz-Signature=..."`
	// Chave do banner no storage (opcional)
	BannerKey *string `json:"bannerKey" example:"groups/123/banner.jpg"`
	// ID do grupo
	GroupId *uint `json:"groupId" example:"123"`
}
