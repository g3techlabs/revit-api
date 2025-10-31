package services

import (
	"github.com/g3techlabs/revit-api/src/core/auth/input"
	"github.com/g3techlabs/revit-api/src/response/generics"
	"github.com/g3techlabs/revit-api/src/utils"
)

func (as *AuthService) ResetPassword(input *input.ResetPassword) error {
	if err := as.validator.Validate(input); err != nil {
		return err
	}

	claims, err := as.tokenService.ValidateResetPassToken(input.ResetToken)
	if err != nil {
		return generics.Unauthorized("Invalid or expired token")
	}

	user, err := as.userRepo.FindUserById(claims.UserID)
	if err != nil {
		return generics.InternalError()
	} else if user == nil {
		return generics.Unauthorized("User does not exist")
	}

	newPassword, err := utils.HashPassword(input.NewPassword)
	if err != nil {
		return generics.InternalError()
	}

	if err = as.userRepo.UpdateUserPassword(claims.UserID, newPassword); err != nil {
		return generics.InternalError()
	}

	return nil
}
