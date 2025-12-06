package response

type GroupMember struct {
	ID            uint    `json:"id"`
	Nickname      string  `json:"nickname"`
	ProfilePicUrl *string `json:"profilePicUrl"`
	Role          string  `json:"role"`
}

type GroupMembersResponse struct {
	Members     []GroupMember `json:"members"`
	TotalPages  uint          `json:"totalPages"`
	CurrentPage uint          `json:"currentPage"`
}
