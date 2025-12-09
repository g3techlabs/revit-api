package response

type FriendshipRequest struct {
	RequesterID uint    `json:"id"`
	Nickname    string  `json:"nickname"`
	ProfilePic  *string `json:"profilePicUrl"`
}

type GetFriendshipRequestsResponse struct {
	Requests    []FriendshipRequest `json:"requests"`
	CurrentPage uint                `json:"currentPage"`
	TotalPages  uint                `json:"totalPages"`
}
