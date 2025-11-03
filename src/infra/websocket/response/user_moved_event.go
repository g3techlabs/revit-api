package response

type UserMovedEvent struct {
	Event   string           `json:"event"`
	Payload UserMovedPayload `json:"payload"`
}

type UserMovedPayload struct {
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
	UserID uint    `json:"userId"`
}
