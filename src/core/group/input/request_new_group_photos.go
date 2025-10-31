package input

type RequestNewGroupPhotos struct {
	MainPhotoContentType *string `json:"mainPhotoContentType" validate:"omitempty,oneof=image/jpeg image/png image/webp"`
	BannerContentType    *string `json:"bannerContentType" validate:"omitempty,oneof=image/jpeg image/png image/webp"`
}
