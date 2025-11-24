package input

// AnswerPendingInvite representa a resposta a um convite pendente de grupo
// @Description Resposta para aceitar ou rejeitar um convite pendente para participar de um grupo
type AnswerPendingInvite struct {
	// Resposta ao convite (accept ou reject)
	Answer string `json:"answer" validate:"required,oneof=accept reject" example:"accept"`
}
