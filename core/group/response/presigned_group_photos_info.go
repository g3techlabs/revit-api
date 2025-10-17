package response

type PresignedGroupPhotosInfo struct {
	PresignedMainPhotoUrl *string `json:"presignedMainPhotoUrl"`
	MainPhotoKey          *string `json:"mainPhotoKey"`
	PresignedBannerUrl    *string `json:"presignedBannerUrl"`
	BannerKey             *string `json:"bannerKey"`
	GroupId               *uint   `json:"groupId"`
}
