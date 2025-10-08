package input

type UpdateProfilePic struct {
	Name string `json:"name" validate:"required"`
}
