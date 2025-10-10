package service

import (
	"github.com/g3techlabs/revit-api/core/users/input"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (us *UserService) ConfirmNewProfilePic(id uint, input *input.ConfirmNewProfilePic) error {
	if err := us.validator.Struct(input); err != nil {
		return err
	}

	err := us.storageService.DoesObjectExist(input.ObjectKey)
	if err != nil {
		return err
	}

	if err := us.userRepo.UpdateUserProfilePic(id, &input.ObjectKey); err != nil {
		us.Log.Errorf("Error updating profile picture for USER ID %d: %s", id, err.Error())
		return generics.InternalError()
	}

	return nil
}
