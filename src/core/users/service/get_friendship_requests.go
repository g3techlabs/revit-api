package service

import (
	"github.com/g3techlabs/revit-api/src/core/users/input"
	"github.com/g3techlabs/revit-api/src/core/users/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (us *UserService) GetFriendshipRequests(userId uint, query *input.GetFriendshipRequestsQuery) (*response.GetFriendshipRequestsResponse, error) {
	if err := us.validator.Validate(query); err != nil {
		return nil, err
	}

	requestsResponse, err := us.userRepo.GetFriendshipRequests(userId, query.Page, query.Limit)
	if err != nil {
		return nil, generics.InternalError()
	}

	return requestsResponse, nil
}
