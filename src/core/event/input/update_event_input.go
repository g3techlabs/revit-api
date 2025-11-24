package input

import "github.com/g3techlabs/revit-api/src/core/event/entities"

// UpdateEventInput representa os dados para atualização de um evento
// @Description Dados opcionais para atualizar informações de um evento existente
type UpdateEventInput struct {
	// Nome do evento (opcional)
	Name *string `json:"name" validate:"omitempty" example:"Novo Nome do Evento"`
	// Descrição do evento (opcional)
	Description *string `json:"description" validate:"omitempty" example:"Nova descrição do evento"`
	// Data e hora do evento no formato RFC3339 (opcional: 2006-01-02T15:04:05Z07:00)
	Date *string `json:"date" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00" example:"2024-12-26T20:00:00-03:00"`
	// Coordenadas geográficas da localização (opcional, requer CityID se fornecido)
	Location *entities.Coordinates `json:"location" validate:"required_with=CityID,omitempty"`
	// ID da cidade (opcional, requer Location se fornecido)
	CityID *uint `json:"cityId" validate:"required_with=Location,omitempty,number,gt=0" example:"2"`
	// ID do grupo associado (opcional)
	GroupID *uint `json:"groupId" validate:"omitempty,number,gt=0" example:"6"`
	// Visibilidade do evento (opcional: public ou private)
	Visibility *string `json:"visibility" validate:"omitempty,oneof=public private" example:"private"`
}
