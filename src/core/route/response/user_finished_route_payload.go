package response

type UserFinishedRouteEvent struct {
	Event   string                   `json:"event"`
	Payload UserFinishedRoutePayload `json:"payload"`
}

type UserFinishedRoutePayload struct {
	UserID       uint   `json:"userId"`
	ArrivalTime  string `json:"arrivalTime"`
	ArrivalOrder int    `json:"arrivalOrder"`
}

func NewUserFinishedRouteEvent(userId uint, arrivalTime string, arrivalOrder int) UserFinishedRouteEvent {
	return UserFinishedRouteEvent{
		Event: "user-finished-route",
		Payload: UserFinishedRoutePayload{
			UserID:       userId,
			ArrivalTime:  arrivalTime,
			ArrivalOrder: arrivalOrder,
		},
	}
}
