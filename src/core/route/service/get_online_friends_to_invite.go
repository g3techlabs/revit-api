package service

import (
	"github.com/g3techlabs/revit-api/src/core/route/response"
	"github.com/g3techlabs/revit-api/src/response/generics"
)

func (s *RouteService) GetOnlineFriendsToInvite(userId uint) (*[]response.OnlineFriendsResponse, error) {
	userFriends, err := s.routeRepo.GetFriendsToInvite(userId)
	if err != nil {
		return nil, generics.InternalError()
	}

	if userFriends == nil {
		return &[]response.OnlineFriendsResponse{}, nil
	}

	// TODO: optimize this shit
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
