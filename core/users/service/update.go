package service

import (
	"time"

	"github.com/g3techlabs/revit-api/core/users/errors"
	"github.com/g3techlabs/revit-api/core/users/input"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (us *UserService) Update(id uint, input *input.UpdateUser) error {
	if err := us.validator.Struct(input); err != nil {
		return err
	}

	birthdate, err := time.Parse("2006-01-02", *input.Birthdate)
	if err != nil {
		return errors.InvalidBirthdateFormat()
	}

	err = us.userRepo.Update(id, &input.Name, &birthdate)
	if err != nil {
		return generics.InternalError()
	}

	return nil
}
