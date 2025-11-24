package input

// PendingInviteAnswer representa a resposta a um convite pendente de evento
// @Description Resposta para aceitar ou rejeitar um convite pendente para participar de um evento
type PendingInviteAnswer struct {
	// Resposta ao convite (accept ou reject)
	Answer string `json:"answer" validate:"required,oneof=accept reject" example:"accept"`
}
