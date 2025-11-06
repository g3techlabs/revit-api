package response

type NearbyUserToRouteResponse struct {
	ID         uint   `json:"id"`
	Nickname   string `json:"nickname"`
	ProfilePic string `json:"profilePic"`
}
