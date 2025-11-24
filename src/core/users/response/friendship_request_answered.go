package response

// FriendshipRequestAnswered representa a resposta após processar uma solicitação de amizade
// @Description Resposta retornada após aceitar ou rejeitar uma solicitação de amizade
type FriendshipRequestAnswered struct {
	// Mensagem descritiva da operação
	Message string `json:"message" example:"Friendship request accepted"`
	// Tipo de operação realizada (accept ou reject)
	Operation string `json:"operation" example:"accept"`
}
