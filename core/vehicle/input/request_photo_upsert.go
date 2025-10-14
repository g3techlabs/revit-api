package input

type RequestPhotoUpsert struct {
	ContentType string `json:"contentType" validate:"required,oneof=image/jpeg image/png image/webp"`
	PhotoType   string `json:"photoType" validate:"required,oneof=main feed"`
}
