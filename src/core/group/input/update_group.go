package input

// UpdateGroup representa os dados para atualização de um grupo
// @Description Dados opcionais para atualizar informações de um grupo existente
type UpdateGroup struct {
	// Nome do grupo (opcional)
	Name *string `json:"name" validate:"omitempty" example:"Novo Nome do Grupo"`
	// Descrição do grupo (opcional)
	Description *string `json:"description" validate:"omitempty" example:"Nova descrição do grupo"`
	// ID da cidade (opcional)
	CityID *uint `json:"cityId" validate:"omitempty,number,gt=0" example:"2"`
	// ID do estado (opcional)
	StateID *uint `json:"stateId" validate:"omitempty,number,gt=0" example:"1"`
	// Visibilidade do grupo (opcional: public ou private)
	Visibility *string `json:"visibility" validate:"omitempty,oneof=public private" example:"private"`
}
