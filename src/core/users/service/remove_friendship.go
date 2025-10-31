package service

import (
	"github.com/g3techlabs/revit-api/src/core/users/errors"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (us *UserService) RemoveFriendship(userId, friendId uint) error {
	if userId == friendId {
		return errors.RequesterAndDestinataryAreTheSame()
	}

	friendUser, err := us.userRepo.FindUserById(friendId)
	if err != nil {
		return generics.InternalError()
	} else if friendUser == nil {
		return errors.UserNotFound("Destinatary was not found")
	}

	if err := us.userRepo.RemoveFriendship(userId, friendId); err != nil {
		if err.Error() == "friendship was not found" {
			return errors.FriendshipNotFound()
		}
		return generics.InternalError()
	}

	return nil
}
