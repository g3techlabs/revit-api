package response

type AdminGroup struct {
	ID   uint   `json:"id" example:"123"`
	Name string `json:"name" example:"Grupo de Ciclistas"`
}

type GetAdminGroupsResponse struct {
	Groups      []AdminGroup `json:"groups"`
	TotalPages  uint         `json:"totalPages" example:"5"`
	CurrentPage uint         `json:"currentPage" example:"1"`
}
