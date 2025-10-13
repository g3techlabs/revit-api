package service

import (
	"github.com/g3techlabs/revit-api/core/users/input"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (us *UserService) ConfirmNewProfilePic(id uint, input *input.ConfirmNewProfilePic) error {
	if err := us.validator.Validate(input); err != nil {
		return err
	}

	err := us.storageService.DoesObjectExist(input.ObjectKey)
	if err != nil {
		return err
	}

	if err := us.userRepo.UpdateUserProfilePic(id, &input.ObjectKey); err != nil {
		return generics.InternalError()
	}

	return nil
}
