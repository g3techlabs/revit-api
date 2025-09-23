package dto

type CreateUser struct {
	Name     string `json:"name" validate:"required"`
	Nickname string `json:"nickname" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,uperandlowerrunes"`
}
