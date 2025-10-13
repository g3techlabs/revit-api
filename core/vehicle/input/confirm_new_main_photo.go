package input

type ConfirmNewMainPhoto struct {
	ObjectKey string `json:"objectKey" validate:"required"`
}
