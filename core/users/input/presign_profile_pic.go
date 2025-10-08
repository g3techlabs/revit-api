package input

type PresignProfilePic struct {
	ContentType string `json:"contentType" validate:"required,oneof=image/jpeg image/png"`
}
