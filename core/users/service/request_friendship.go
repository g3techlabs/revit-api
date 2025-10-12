package service

import (
	"github.com/g3techlabs/revit-api/core/users/errors"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (us *UserService) RequestFriendship(userId, destinataryId uint) error {
	if userId == destinataryId {
		return errors.RequesterAndDestinataryAreTheSame()
	}

	destinataryUser, err := us.userRepo.FindUserById(destinataryId)
	if err != nil {
		return generics.InternalError()
	} else if destinataryUser == nil {
		return errors.UserNotFound("Destinatary was not found")
	}

	areFriends, err := us.userRepo.AreFriends(userId, destinataryId)
	if err != nil {
		return generics.InternalError()
	}
	if areFriends {
		return errors.FriendsAlready()
	}

	if err := us.userRepo.RequestFriendship(userId, destinataryId); err != nil {
		return generics.InternalError()
	}

	return nil
}
