package service

import (
	"github.com/g3techlabs/revit-api/src/core/users/errors"
	"github.com/g3techlabs/revit-api/src/core/users/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (us *UserService) GetUser(requesterId, userId uint) (*response.GetUserResponse, error) {
	user, err := us.userRepo.GetUserDetails(userId)
	if err != nil {
		return nil, generics.InternalError()
	} else if user == nil {
		return nil, errors.UserNotFound("User not found")
	}

	if requesterId != userId {
		isFriend, err := us.userRepo.AreFriendsAccepted(requesterId, userId)
		if err != nil {
			return nil, generics.InternalError()
		}
		user.IsFriend = isFriend
	} else {
		user.IsFriend = false
	}

	return user, nil
}
