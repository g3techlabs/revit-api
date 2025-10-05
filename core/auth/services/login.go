package services

import (
	"net/mail"

	"github.com/g3techlabs/revit-api/core/auth/input"
	response "github.com/g3techlabs/revit-api/core/auth/response"
	"github.com/g3techlabs/revit-api/core/users/models"
	"github.com/g3techlabs/revit-api/response/generics"
	"github.com/g3techlabs/revit-api/utils"
)

func (as *AuthService) Login(loginCredentials *input.LoginCredentials) (*response.AuthTokensResponse, error) {
	if err := as.validator.Struct(loginCredentials); err != nil {
		return nil, err
	}

	user, err := as.findUserByIdentifier(loginCredentials.Identifier)
	if err != nil {
		return nil, generics.InternalError()
	} else if user == nil {
		return nil, generics.Unauthorized("Invalid credentials")
	}

	if !utils.CheckPasswordHash(loginCredentials.Password, user.Password) {
		return nil, generics.Unauthorized("Invalid credentials")
	}

	accessToken, refreshToken, err := as.tokenService.GenerateAuthTokens(user.ID)
	if err != nil {
		return nil, generics.InternalError()
	}

	return &response.AuthTokensResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func isAnEmail(value string) bool {
	_, err := mail.ParseAddress(value)
	return err == nil
}

func (as *AuthService) findUserByIdentifier(identifier string) (*models.User, error) {
	isIdentifierAnEmail := isAnEmail(identifier)

	user, err := new(models.User), *new(error)
	if isIdentifierAnEmail {
		user, err = as.userRepo.FindUserByEmail(identifier)

	} else {
		user, err = as.userRepo.FindUserByNickname(identifier)
	}

	return user, err
}
