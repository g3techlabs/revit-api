package response

type GetPendingInvites struct {
	EventID    uint   `json:"eventId"`
	EventName  string `json:"eventName"`
	EventPhoto string `json:"eventPhoto"`
	InvitedBy  string `json:"invitedBy"`
}
