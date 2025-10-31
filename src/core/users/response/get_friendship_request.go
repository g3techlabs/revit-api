package response

type FriendshipRequest struct {
	RequesterID uint    `json:"id"`
	Nickname    string  `json:"nickname"`
	ProfilePic  *string `json:"ProfilePicUrl"`
}
