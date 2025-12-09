package input

type RequestNewPhotoInput struct {
	PhotoContentType string `json:"photoContentType" validate:"required,oneof=image/jpeg image/png image/webp"`
}
