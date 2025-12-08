package response

// GetPendingInvitesResponse representa um convite pendente de evento
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

// GetPendingInvitesPaginatedResponse representa a resposta paginada de convites pendentes
// @Description Resposta com lista de convites pendentes, página atual e total de páginas
type GetPendingInvitesPaginatedResponse struct {
	Invites     []GetPendingInvitesResponse `json:"invites"`
	CurrentPage uint                        `json:"currentPage"`
	TotalPages  uint                        `json:"totalPages"`
}
