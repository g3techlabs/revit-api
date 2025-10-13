package services

import (
	"fmt"
	"strings"

	"github.com/g3techlabs/revit-api/config"
	"github.com/g3techlabs/revit-api/core/auth/input"
	"github.com/g3techlabs/revit-api/response/generics"
)

var resetTokenExpirationInMinutes int = config.GetIntVariable("RESET_TOKEN_EXPIRATION")
var appName = config.Get("APP_NAME")

func (as *AuthService) SendPassResetEmail(input *input.Identifier) error {
	if err := as.validator.Validate(input); err != nil {
		return err
	}

	user, err := as.findUserByIdentifier(input.Identifier, input.IdentifierType)
	if err != nil {
		return generics.InternalError()
	} else if user == nil {
		return nil
	}

	deepLink, err := as.generateDeepLink(user.ID)
	if err != nil {
		return generics.InternalError()
	}

	if err := as.emailService.SendPassResetEmail(user.Email, user.Name, deepLink, resetTokenExpirationInMinutes); err != nil {
		return generics.InternalError()
	}

	return nil
}

func (as *AuthService) generateDeepLink(userId uint) (string, error) {
	resetPassToken, err := as.tokenService.GenerateResetPassToken(userId)
	if err != nil {
		return "", nil
	}

	deepLink := fmt.Sprintf("%s://reset_password?t=%s", strings.ToLower(appName), resetPassToken)

	return deepLink, nil
}
