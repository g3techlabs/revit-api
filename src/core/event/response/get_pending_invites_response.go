package response

type GetPendingInvitesResponse struct {
	EventID    uint   `json:"eventId"`
	EventName  string `json:"eventName"`
	EventPhoto string `json:"eventPhoto"`
	InvitedBy  string `json:"invitedBy"`
}

type GetPendingInvitesPaginatedResponse struct {
	Invites     []GetPendingInvitesResponse `json:"invites"`
	CurrentPage uint                        `json:"currentPage"`
	TotalPages  uint                        `json:"totalPages"`
}
