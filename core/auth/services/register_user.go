package services

import (
	"fmt"

	"github.com/g3techlabs/revit-api/core/auth/errors"
	usersInput "github.com/g3techlabs/revit-api/core/users/input"
	usersResponse "github.com/g3techlabs/revit-api/core/users/response"
	"github.com/g3techlabs/revit-api/utils"
	"github.com/g3techlabs/revit-api/utils/generics"
)

func (as *AuthService) RegisterUser(input usersInput.CreateUser) (*usersResponse.UserCreatedResponse, error) {
	conflictErrs := make(map[string]string)
	nicknameErr, emailErr := as.verifyUniqueFieldsAvailability(input.Nickname, input.Email)
	if nicknameErr != nil {
		conflictErrs["nickname"] = nicknameErr.Error()
	}
	if emailErr != nil {
		conflictErrs["email"] = emailErr.Error()
	}
	if len(conflictErrs) > 0 {
		return nil, errors.NewConflictErrors(conflictErrs)
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

func (as AuthService) verifyUniqueFieldsAvailability(nickname string, email string) (error, error) {
	var emailError error
	var nicknameError error

	if !as.isNicknameAvailable(nickname) {
		nicknameError = fmt.Errorf("nickname already in use")
	}

	if !as.isEmailAvailable(email) {
		emailError = fmt.Errorf("email already in use")
	}

	return nicknameError, emailError
}

func (as AuthService) isNicknameAvailable(nickname string) bool {
	foundUser, _ := as.userRepo.FindUserByNickname(nickname)
	return foundUser == nil
}

func (as AuthService) isEmailAvailable(email string) bool {
	foundUser, _ := as.userRepo.FindUserByEmail(email)
	return foundUser == nil
}
