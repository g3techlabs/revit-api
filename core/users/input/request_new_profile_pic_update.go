package input

type RequestProfilePicUpdate struct {
	ContentType string `json:"contentType" validate:"required,oneof=image/jpeg image/png"`
}
