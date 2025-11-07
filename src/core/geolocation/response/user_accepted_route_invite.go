package response

type UserAcceptedRouteInvite struct {
	Event   string      `json:"event"`
	Payload UserDetails `json:"userDetails"`
}

type UserDetails struct {
	UserId     uint   `json:"userId"`
	Nickname   string `json:"nickname"`
	ProfilePic string `json:"profilePic"`
}

func NewUserAcceptedRouteInviteEvent(userDetails UserDetails) *UserAcceptedRouteInvite {
	return &UserAcceptedRouteInvite{
		Event:   "user-accepted-route-invite",
		Payload: userDetails,
	}
}
