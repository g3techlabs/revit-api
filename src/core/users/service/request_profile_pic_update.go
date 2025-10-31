package service

import (
	"fmt"

	"github.com/g3techlabs/revit-api/src/core/users/input"
	"github.com/g3techlabs/revit-api/src/core/users/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (us *UserService) RequestProfilePicUpdate(userId uint, input *input.RequestProfilePicUpdate) (*response.ProfilePicPresignedURL, error) {
	if err := us.validator.Validate(input); err != nil {
		return nil, err
	}

	profilePicKey := fmt.Sprintf("users/%d/profile", userId)
	presignedUrl, err := us.storageService.PresignPutObjectURL(profilePicKey, input.ContentType)
	if err != nil {
		return nil, generics.InternalError()
	}

	return &response.ProfilePicPresignedURL{PresignedURL: presignedUrl, ObjectKey: profilePicKey}, nil
}
