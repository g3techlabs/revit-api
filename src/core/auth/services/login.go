package services

import (
	"net/mail"

	"github.com/g3techlabs/revit-api/src/core/auth/errors"
	"github.com/g3techlabs/revit-api/src/core/auth/input"
	response "github.com/g3techlabs/revit-api/src/core/auth/response"
	"github.com/g3techlabs/revit-api/src/db/models"
	"github.com/g3techlabs/revit-api/src/response/generics"
	"github.com/g3techlabs/revit-api/src/utils"
)

func (as *AuthService) Login(loginCredentials *input.LoginCredentials) (*response.AuthTokensResponse, error) {
	if err := as.validator.Validate(loginCredentials); err != nil {
		return nil, err
	}

	user, err := as.findUserByIdentifier(loginCredentials.Identifier, loginCredentials.IdentifierType)
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

	var profilePicUrl *string
	if user.ProfilePic != nil {
		profilePicUrl = utils.MountCloudFrontUrl(*user.ProfilePic)
	}

	return &response.AuthTokensResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ID:           user.ID,
		ProfilePicUrl: profilePicUrl,
		Name:         user.Name,
		Nickname:     user.Nickname,
	}, nil
}

func isAnEmail(value string) bool {
	_, err := mail.ParseAddress(value)
	return err == nil
}

func (as *AuthService) findUserByIdentifier(identifier string, identifierType string) (*models.User, error) {
	user, err := new(models.User), *new(error)

	if identifierType == "email" {
		if !isAnEmail(identifier) {
			return nil, errors.NotAnEmail()
		}

		user, err = as.userRepo.FindUserByEmail(identifier)
		if err != nil {
			return nil, err
		}

		return user, nil
	} else {
		user, err = as.userRepo.FindUserByNickname(identifier)
	}

	return user, err
}
