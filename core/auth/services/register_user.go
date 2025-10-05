package services

import (
	"github.com/g3techlabs/revit-api/core/auth/errors"
	usersInput "github.com/g3techlabs/revit-api/core/users/input"
	usersResponse "github.com/g3techlabs/revit-api/core/users/response"
	"github.com/g3techlabs/revit-api/response/generics"
	"github.com/g3techlabs/revit-api/utils"
)

func (as *AuthService) RegisterUser(input *usersInput.CreateUser) (*usersResponse.UserCreatedResponse, error) {
	if err := as.validator.Struct(input); err != nil {
		return nil, err
	}

	nicknameTaken, emailTaken := as.verifyUniqueFieldsAvailability(input.Nickname, input.Email)
	if nicknameTaken || emailTaken {
		return nil, errors.NewConflictError(emailTaken, nicknameTaken)
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, generics.InternalError()
	}
	input.Password = hashedPassword

	user := input.ToUserModel()
	err = as.userRepo.RegisterUser(user)
	if err != nil {
		return nil, generics.InternalError()
	}

	userResponse := user.ToUserCreatedResponse()

	return userResponse, nil
}

func (as AuthService) verifyUniqueFieldsAvailability(nickname string, email string) (bool, bool) {
	return as.isNicknameTaken(nickname), as.isEmailTaken(email)
}

func (as AuthService) isNicknameTaken(nickname string) bool {
	foundUser, _ := as.userRepo.FindUserByNickname(nickname)
	return foundUser != nil
}

func (as AuthService) isEmailTaken(email string) bool {
	foundUser, _ := as.userRepo.FindUserByEmail(email)
	return foundUser != nil
}
