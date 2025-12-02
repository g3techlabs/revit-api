package response

import (
	"time"

	"gorm.io/datatypes"
)

// GetGroupsResponse representa um grupo retornado na listagem
// @Description Informações completas de um grupo retornado na busca/listagem
type GroupResponse struct {
	// ID do grupo
	ID uint `json:"id" example:"123"`
	// Nome do grupo
	Name string `json:"name" example:"Grupo de Ciclistas"`
	// Descrição do grupo
	Description string `json:"description" example:"Grupo para ciclistas da cidade"`
	// URL da foto principal do grupo (opcional)
	MainPhoto *string `json:"mainPhoto" example:"https://example.com/groups/123/main.jpg"`
	// URL do banner do grupo (opcional)
	Banner *string `json:"banner" example:"https://example.com/groups/123/banner.jpg"`
	// Data de criação do grupo
	CreatedAt time.Time `json:"createdAt" example:"2024-01-15T10:30:00Z"`
	// Visibilidade do grupo (public ou private)
	Visibility string `json:"visibility" example:"public"`
	// Nome da cidade
	City string `json:"city" example:"São Paulo"`
	// Nome do estado
	State string `json:"state" example:"SP"`
	// Tipo de membro do usuário no grupo (opcional: owner, admin, member)
	MemberType *string `json:"memberType" example:"member"`
	// Lista de amigos no grupo (formato JSON)
	FriendsInGroup datatypes.JSON `json:"friendsInGroup"`
	// Total de membros no grupo
	TotalMembers uint `json:"totalMembers" example:"150"`
}

type GetGroupsResponse struct {
	Groups      []GroupResponse `json:"groups"`
	CurrentPage uint            `json:"currentPage" example:"1"`
	TotalPages  uint            `json:"totalPages" example:"10"`
}
