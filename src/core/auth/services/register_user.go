package services

import (
	"strings"

	"github.com/g3techlabs/revit-api/src/core/auth/errors"
	usersInput "github.com/g3techlabs/revit-api/src/core/users/input"
	usersResponse "github.com/g3techlabs/revit-api/src/core/users/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
	"github.com/g3techlabs/revit-api/src/utils"
)

func (as *AuthService) RegisterUser(input *usersInput.CreateUser) (*usersResponse.UserCreatedResponse, error) {
	if err := as.validator.Validate(input); err != nil {
		return nil, err
	}

	nicknameTaken, emailTaken := as.verifyUniqueFieldsAvailability(input.Nickname, input.Email)
	if nicknameTaken || emailTaken {
		return nil, errors.NewConflictError(emailTaken, nicknameTaken)
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		as.log.Error("Error hashing password: ", err.Error())
		return nil, generics.InternalError()
	}
	input.Password = hashedPassword

	input.Name, input.Nickname = strings.ToLower(input.Name), strings.ToLower(input.Nickname)

	user, err := input.ToUserModel()
	if err != nil {
		as.log.Error("Error parsing birthdate from CreateUser to UserModel: ", err.Error())
		return nil, generics.InternalError()
	}

	err = as.userRepo.RegisterUser(user)
	if err != nil {
		as.log.Error("Error registering user: ", err.Error())
		return nil, generics.InternalError()
	}

	userResponse := user.ToUserCreatedResponse()

	return userResponse, nil
}

func (as *AuthService) verifyUniqueFieldsAvailability(nickname string, email string) (bool, bool) {
	isEmailTaken, _ := as.isEmailTaken(email)
	isNicknameTaken, _ := as.isNicknameTaken(nickname)

	return isNicknameTaken, isEmailTaken
}
