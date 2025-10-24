package service

import (
	"github.com/g3techlabs/revit-api/core/users/input"
	"github.com/g3techlabs/revit-api/core/users/response"
	"github.com/g3techlabs/revit-api/response/generics"
)

func (us *UserService) GetFriendshipRequests(userId uint, query *input.GetFriendshipRequestsQuery) (*[]response.FriendshipRequest, error) {
	if err := us.validator.Validate(query); err != nil {
		return nil, err
	}

	requests, err := us.userRepo.GetFriendshipRequests(userId, query.Page, query.Limit)
	if err != nil {
		return nil, generics.InternalError()
	}

	return requests, nil
}
