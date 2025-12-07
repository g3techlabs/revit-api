package input

import (
	"time"

	"github.com/g3techlabs/revit-api/src/db/models"
)

type CreateUser struct {
	Name      string  `json:"name" validate:"required" example:"João Silva"`
	Nickname  string  `json:"nickname" validate:"required,min=3,max=32,lowercase" example:"joaosilva"`
	Email     string  `json:"email" validate:"required,email" example:"joao@email.com"`
	Password  string  `json:"password" validate:"required,password" example:"senhaSegura123"`
	Birthdate *string `json:"birthdate" validate:"omitempty,datetime=2006-01-02" example:"1990-01-15"`
}

func (input CreateUser) ToUserModel() (*models.User, error) {
	var birthdate *time.Time

	if input.Birthdate != nil && *input.Birthdate != "" {
		t, err := time.Parse("2006-01-02", *input.Birthdate)
		if err != nil {
			return nil, err
		}
		birthdate = &t
	}

	return &models.User{
		Name:      input.Name,
		Email:     input.Email,
		Nickname:  input.Nickname,
		Password:  input.Password,
		Birthdate: birthdate,
	}, nil
}
