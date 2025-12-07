package response

type EventSubscriber struct {
	ID            uint    `json:"id"`
	Nickname      string  `json:"nickname"`
	ProfilePicUrl *string `json:"profilePicUrl"`
	Role          string  `json:"role"`
}

type EventSubscribersResponse struct {
	Subscribers []EventSubscriber `json:"subscribers"`
	TotalPages  uint              `json:"totalPages"`
	CurrentPage uint              `json:"currentPage"`
}
