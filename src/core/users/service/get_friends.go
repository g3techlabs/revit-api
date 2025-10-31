package service

import (
	"github.com/g3techlabs/revit-api/src/core/users/input"
	"github.com/g3techlabs/revit-api/src/core/users/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
	"github.com/g3techlabs/revit-api/src/utils"
)

func (us *UserService) GetFriends(userId uint, params *input.GetUsersQuery) (*[]response.Friend, error) {
	if err := us.validator.Validate(us); err != nil {
		return nil, err
	}

	friends, err := us.userRepo.GetFriends(userId, params.Page, params.Limit, params.Nickname)
	if err != nil {
		return nil, generics.InternalError()
	}

	response := make([]response.Friend, 0, len(*friends))
	for i := range *friends {
		friend := (*friends)[i]
		if friend.ProfilePicUrl != nil {
			friend.ProfilePicUrl = utils.MountCloudFrontUrl(*friend.ProfilePicUrl)
		}
		response = append(response, friend)
	}

	return &response, nil
}
