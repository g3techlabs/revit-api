package service

import (
	"strings"

	"github.com/g3techlabs/revit-api/config"
	"github.com/g3techlabs/revit-api/core/users/input"
	"github.com/g3techlabs/revit-api/core/users/response"
	"github.com/g3techlabs/revit-api/response/generics"
)

var cloudFrontUrl = config.Get("AWS_CLOUDFRONT_URL")

func (us *UserService) GetUsers(params *input.GetUsersQuery) (*[]response.GetUserResponse, error) {
	if err := us.validator.Validate(params); err != nil {
		return nil, err
	}

	users, err := us.userRepo.GetUsers(params.Page, params.Limit, params.Nickname)
	if err != nil {
		return nil, generics.InternalError()
	}

	response := make([]response.GetUserResponse, 0, len(*users))
	for i := range *users {
		user := (*users)[i]
		if user.ProfilePic != nil {
			user.ProfilePic = us.mountProfilePicURL(*user.ProfilePic)
		}
		response = append(response, *user.ToGetUserResponse())
	}

	return &response, nil
}

func (us *UserService) mountProfilePicURL(objectKey string) *string {
	if !strings.HasSuffix(cloudFrontUrl, "/") {
		cloudFrontUrl += "/"
	}
	mountedUrl := cloudFrontUrl + objectKey
	return &mountedUrl
}
