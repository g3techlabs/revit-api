package service

import (
	"github.com/g3techlabs/revit-api/core/users/errors"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (us *UserService) AcceptFriendshipRequest(userId, requesterId uint) error {
	if userId == requesterId {
		return errors.RequesterSameAsReceiver()
	}

	requester, err := us.userRepo.FindUserById(requesterId)
	if err != nil {
		return generics.InternalError()
	} else if requester == nil {
		return errors.UserNotFound("Requester was not found")
	}

	if err := us.userRepo.AcceptFriendshipRequest(userId, requesterId); err != nil {
		if err.Error() == "friendship request was not found" {
			return errors.FriendshipRequestNotFound()
		}
		return generics.InternalError()
	}

	return nil
}
