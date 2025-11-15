package service

import (
	"github.com/g3techlabs/revit-api/src/core/users/input"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (us *UserService) CheckIfNicknameAvailable(nickname *input.NicknameInput) (bool, error) {
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

func (us *UserService) isNicknameTaken(nickname string) (bool, error) {
	foundUser, err := us.userRepo.FindUserByNickname(nickname)
	return foundUser != nil, err
}
