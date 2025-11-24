package response

// GetPendingInvites representa um convite pendente de evento
// @Description Informações sobre um convite pendente para participar de um evento
type GetPendingInvitesResponse struct {
	// ID do evento
	EventID uint `json:"eventId" example:"456"`
	// Nome do evento
	EventName string `json:"eventName" example:"Pedal Noturno"`
	// URL da foto do evento
	EventPhoto string `json:"eventPhoto" example:"https://example.com/events/456/photo.jpg"`
	// Apelido do usuário que enviou o convite
	InvitedBy string `json:"invitedBy" example:"joaosilva"`
}
