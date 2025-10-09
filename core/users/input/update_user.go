package input

type UpdateUser struct {
	Name             string  `json:"name" validate:"omitempty"`
	Birthdate        *string `json:"birthdate" validate:"omitempty,datetime=2006-01-02"`
	RemoveProfilePic *bool   `json:"removeProfilePic" validate:"omitempty"`
}
