package service

import (
	"time"

	"github.com/g3techlabs/revit-api/core/users/errors"
	"github.com/g3techlabs/revit-api/core/users/input"
	"github.com/g3techlabs/revit-api/response/generics"
)

const ISO_8601 string = "2006-01-02"

func (us *UserService) Update(id uint, input *input.UpdateUser) error {
	if err := us.validator.Validate(input); err != nil {
		return err
	}

	birthdate, err := time.Parse(ISO_8601, *input.Birthdate)
	if err != nil {
		return errors.InvalidBirthdateFormat()
	}

	if input.RemoveProfilePic != nil && *input.RemoveProfilePic {
		user, err := us.userRepo.FindUserById(id)
		if err != nil {
			return generics.InternalError()
		}
		if user.ProfilePic != nil {
			if err := us.removeUserProfilePic(user.ID, *user.ProfilePic); err != nil {
				return err
			}
		}
	}

	err = us.userRepo.Update(id, &input.Name, &birthdate)
	if err != nil {
		return generics.InternalError()
	}

	return nil
}

func (us *UserService) removeUserProfilePic(userId uint, objectKey string) error {
	if err := us.userRepo.UpdateUserProfilePic(userId, nil); err != nil {
		return generics.InternalError()
	}

	if err := us.storageService.DeleteObject(objectKey); err != nil {
		return generics.InternalError()
	}

	return nil
}
