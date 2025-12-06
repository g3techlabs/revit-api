package response

type SimpleGroup struct {
	ID           uint    `json:"id" example:"123"`
	Name         string  `json:"name" example:"Grupo de Ciclistas"`
	MainPhoto    *string `json:"mainPhoto"`
	MemberType   string  `json:"memberType"`
	Banner       *string `json:"banner"`
	TotalMembers uint    `json:"totalMembers"`
}

type GetGroupsResponse struct {
	Groups      []SimpleGroup `json:"groups"`
	CurrentPage uint          `json:"currentPage"`
	TotalPages  uint          `json:"totalPages"`
}
