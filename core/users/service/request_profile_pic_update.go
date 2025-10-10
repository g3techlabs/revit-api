package service

import (
	"fmt"

	"github.com/g3techlabs/revit-api/core/users/errors"
	"github.com/g3techlabs/revit-api/core/users/input"
	"github.com/g3techlabs/revit-api/core/users/response"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (us *UserService) RequestProfilePicUpdate(userId uint, input *input.RequestProfilePicUpdate) (*response.ProfilePicPresignedURL, error) {
	if err := us.validator.Struct(input); err != nil {
		return nil, err
	}

	extension := us.mapContentTypeToExtension(input.ContentType)
	if extension == "" {
		return nil, errors.InvalidFileExtensionError()
	}
	profilePicKey := fmt.Sprintf("users/%d/profile%s", userId, extension)

	presignedUrl, err := us.storageService.PresignPutObjectURL(profilePicKey, input.ContentType)
	if err != nil {
		us.Log.Errorf("Error generating presigned URL for USER ID %d profile picture: %s", userId, err.Error())
		return nil, generics.InternalError()
	}

	return &response.ProfilePicPresignedURL{PresignedURL: presignedUrl, ObjectKey: profilePicKey}, nil
}

func (us *UserService) mapContentTypeToExtension(contentType string) string {
	switch contentType {
	case "image/jpeg":
		return ".jpg"
	case "image/png":
		return ".png"
	default:
		return ""
	}
}
