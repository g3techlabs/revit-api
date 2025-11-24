package services

import (
	"github.com/g3techlabs/revit-api/src/core/auth/input"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (as *AuthService) CheckIfEmailAvailable(email *input.EmailInput) (bool, error) {
	if err := as.validator.Validate(email); err != nil {
		return false, err
	}

	isEmailTaken, err := as.isEmailTaken(email.Email)
	if err != nil {
		return false, generics.InternalError()
	}
	if isEmailTaken {
		return false, nil
	}

	return true, nil
}

func (as *AuthService) isEmailTaken(email string) (bool, error) {
	foundUser, err := as.userRepo.FindUserByEmail(email)
	return foundUser != nil, err
}
