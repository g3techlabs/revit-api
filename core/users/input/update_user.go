package input

import "mime/multipart"

type UpdateUser struct {
	Name       string                `json:"name" validate:"omitempty"`
	Birthdate  *string               `json:"birthdate" validate:"omitempty,datetime=2006-01-02"`
	ProfilePic *multipart.FileHeader `validate:"omitempty,profilepic"`
}
