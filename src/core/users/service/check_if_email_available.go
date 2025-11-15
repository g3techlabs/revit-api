package service

import (
	"github.com/g3techlabs/revit-api/src/core/users/input"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (us *UserService) CheckIfEmailAvailable(email *input.EmailInput) (bool, error) {
	if err := us.validator.Validate(email); err != nil {
		return false, err
	}

	isEmailTaken, err := us.isEmailTaken(email.Email)
	if err != nil {
		return false, generics.InternalError()
	}
	if isEmailTaken {
		return false, nil
	}

	return true, nil
}

func (us *UserService) isEmailTaken(email string) (bool, error) {
	foundUser, err := us.userRepo.FindUserByEmail(email)
	return foundUser != nil, err
}
