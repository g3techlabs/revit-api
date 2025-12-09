package input

type GetAdminEventsInput struct {
	Name  string `json:"name" validate:"omitempty"`
	Limit uint   `json:"limit" validate:"required,number,gt=0"`
	Page  uint   `json:"page" validate:"required,number,gt=0"`
}
