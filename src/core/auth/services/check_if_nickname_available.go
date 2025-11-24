package services

import (
	"github.com/g3techlabs/revit-api/src/core/auth/input"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (us *AuthService) CheckIfNicknameAvailable(nickname *input.NicknameInput) (bool, error) {
	if err := us.validator.Validate(nickname); err != nil {
		return false, err
	}

	isNicknameTaken, err := us.isNicknameTaken(nickname.Nickname)
	if err != nil {
		return false, generics.InternalError()
	}
	if isNicknameTaken {
		return false, nil
	}

	return true, nil
}

func (us *AuthService) isNicknameTaken(nickname string) (bool, error) {
	foundUser, err := us.userRepo.FindUserByNickname(nickname)
	return foundUser != nil, err
}
