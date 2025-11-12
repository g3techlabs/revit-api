package response

type UserFinishedRoutePayload struct {
	UserID       uint   `json:"userId"`
	ArrivalTime  string `json:"arrivalTime"`
	ArrivalOrder int    `json:"arrivalOrder"`
}

func NewUserFinishedRoutePayload(userId uint, arrivalTime string, arrivalOrder int) UserFinishedRoutePayload {
	return UserFinishedRoutePayload{
		UserID:       userId,
		ArrivalTime:  arrivalTime,
		ArrivalOrder: arrivalOrder,
	}
}
