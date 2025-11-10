package service

import (
	"github.com/g3techlabs/revit-api/src/core/route/input"
	"github.com/g3techlabs/revit-api/src/core/route/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (s *RouteService) GetOnlineFriendsToInvite(userId uint, query *input.GetOnlineFriendsToInviteQuery) (*[]response.OnlineFriendsResponse, error) {
	if err := s.validator.Validate(query); err != nil {
		return nil, err
	}

	userFriends, err := s.routeRepo.GetFriendsToInvite(userId, query.Page, query.Limit)
	if err != nil {
		return nil, generics.InternalError()
	}

	if userFriends == nil {
		return &[]response.OnlineFriendsResponse{}, nil
	}

	friendIDs := make([]uint, len(*userFriends))
	for i, friend := range *userFriends {
		friendIDs[i] = friend.FriendId
	}

	onlineStatuses, err := s.geoLocationService.CheckUsersAreOnline(friendIDs)
	if err != nil {
		return nil, generics.InternalError()
	}

	onlineFriends := []response.OnlineFriendsResponse{}
	for i, friend := range *userFriends {
		if onlineStatuses[i] {
			onlineFriends = append(onlineFriends, friend)
		}
	}

	return &onlineFriends, nil
}
