package input

import "github.com/g3techlabs/revit-api/core/users/models"

type CreateUser struct {
	Name     string `json:"name" validate:"required"`
	Nickname string `json:"nickname" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,uperandlowerrunes"`
}

func (input CreateUser) ToUserModel() *models.User {
	return &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Nickname: input.Nickname,
		Password: input.Password,
	}
}
