package input

type RequestMainPhotoUpdate struct {
	ContentType string `json:"contentType" validate:"required,oneof=image/jpeg image/png image/webp"`
}
