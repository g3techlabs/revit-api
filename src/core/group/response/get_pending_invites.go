package response

type GetPendingInvites struct {
	GroupID        uint    `json:"groupId" example:"123"`
	GroupName      string  `json:"groupName" example:"Honda Club"`
	GroupMainPhoto *string `json:"groupMainPhoto" example:"https://example.com/groups/123/main.jpg"`
	InvitedBy      string  `json:"invitedBy" example:"hondeiro2000"`
}

type GetPendingInvitesPaginatedResponse struct {
	Invites     []GetPendingInvites `json:"invites"`
	CurrentPage uint                `json:"currentPage"`
	TotalPages  uint                `json:"totalPages"`
}
