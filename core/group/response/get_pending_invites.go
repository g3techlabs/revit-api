package response

type GetPendingInvites struct {
	GroupName      string  `json:"groupName"`
	GroupMainPhoto *string `json:"groupMainPhoto"`
	InvitedBy      string  `json:"invitedBy"`
}
