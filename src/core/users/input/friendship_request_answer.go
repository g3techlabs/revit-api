package input

// FriendshipRequestAnswer representa a resposta a uma solicitação de amizade
// @Description Resposta para aceitar ou rejeitar uma solicitação de amizade pendente
type FriendshipRequestAnswer struct {
	// Resposta à solicitação (accept ou reject)
	Answer string `json:"answer" validate:"required,oneof=accept reject" example:"accept"`
}
