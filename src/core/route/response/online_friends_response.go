package response

type OnlineFriendsResponse struct {
	FriendId   uint   `json:"friendId"`
	Nickname   string `json:"nickname"`
	ProfilePic string `json:"profilePic"`
}
