package service

import (
	"github.com/g3techlabs/revit-api/src/core/users/errors"
	"github.com/g3techlabs/revit-api/src/core/users/input"
	"github.com/g3techlabs/revit-api/src/core/users/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (us *UserService) AnswerFriendshipRequest(userId, requesterId uint, answer *input.FriendshipRequestAnswer) (*response.FriendshipRequestAnswered, error) {
	if err := us.validator.Validate(answer); err != nil {
		return nil, err
	}

	if userId == requesterId {
		return nil, errors.RequesterAndDestinataryAreTheSame()
	}

	if exists, err := us.verifyIfRequesterExists(requesterId); err != nil {
		return nil, generics.InternalError()
	} else if !exists {
		return nil, errors.UserNotFound("Requester was not found")
	}

	if err := us.respondBasedOnAnswer(userId, requesterId, answer.Answer); err != nil {
		if err.Error() == "friendship request was not found" {
			return nil, errors.FriendshipRequestNotFound()
		} else {
			return nil, generics.InternalError()
		}
	}

	return &response.FriendshipRequestAnswered{
		Message: "Request properly answered", Operation: answer.Answer,
	}, nil
}

func (us *UserService) verifyIfRequesterExists(requesterId uint) (bool, error) {
	requester, err := us.userRepo.FindUserById(requesterId)
	if err != nil {
		return false, err
	} else if requester == nil {
		return false, nil
	}

	return true, nil
}

func (us *UserService) respondBasedOnAnswer(userId, requesterId uint, answer string) error {
	if answer == "accept" {
		return us.userRepo.AcceptFriendshipRequest(userId, requesterId)
	} else {
		return us.userRepo.RejectFriendshipRequest(userId, requesterId)
	}
}
