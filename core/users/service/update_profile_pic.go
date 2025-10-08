package service

import (
	"github.com/g3techlabs/revit-api/core/users/input"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (us *UserService) UpdateProfilePic(id uint, input *input.UpdateProfilePic) error {
	if err := us.validator.Struct(input); err != nil {
		return err
	}
	err := us.storageService.DoesObjectExist(input.Name)
	if err != nil {
		return err
	}

	if err := us.userRepo.UpdateUserProfilePic(id, input.Name); err != nil {
		// TODO: log error
		return generics.InternalError()
	}

	return nil
}
