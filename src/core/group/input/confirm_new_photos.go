package input

type ConfirmNewPhotos struct {
	MainPhotoKey *string `json:"mainPhotoKey" validate:"omitempty"`
	BannerKey    *string `json:"bannerKey" validate:"omitempty"`
}
