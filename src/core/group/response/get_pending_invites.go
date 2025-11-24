package response

// GetPendingInvites representa um convite pendente de grupo
// @Description Informações sobre um convite pendente para participar de um grupo
type GetPendingInvites struct {
	// Nome do grupo
	GroupName string `json:"groupName" example:"Honda Club"`
	// URL da foto principal do grupo (se houver)
	GroupMainPhoto *string `json:"groupMainPhoto" example:"https://example.com/groups/123/main.jpg"`
	// Apelido do usuário que enviou o convite
	InvitedBy string `json:"invitedBy" example:"hondeiro2000"`
}
