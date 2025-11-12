package response

type ParticipantLeftRoutePayload struct {
	UserID uint `json:"userId"`
}

func NewParticipantLeftRoutePayload(userId uint) ParticipantLeftRoutePayload {
	return ParticipantLeftRoutePayload{
		UserID: userId,
	}
}
