package response

import (
	"time"

	"gorm.io/datatypes"
)

// GetEventResponse representa um evento retornado na listagem ou detalhamento
// @Description Informações completas de um evento retornado na busca/listagem
type GetEventResponse struct {
	// ID do evento
	ID uint `json:"id" example:"456"`
	// Nome do evento
	Name string `json:"name" example:"Pedal Noturno"`
	// Descrição do evento
	Description string `json:"description" example:"Passeio ciclístico noturno pela cidade"`
	// Data e hora do evento
	Date time.Time `json:"date" example:"2024-12-25T20:00:00Z"`
	// Visibilidade do evento (public ou private)
	Visibility string `json:"visibility" example:"public"`
	// URL da foto do evento (opcional)
	Photo *string `json:"photo" example:"https://example.com/events/456/photo.jpg"`
	// Número de inscritos no evento (opcional)
	SubscribersCount *uint `json:"subscribersCount" example:"25"`
	// Endereço do evento (formato JSON)
	Address datatypes.JSON `json:"address" swaggertype:"object"`
	// Coordenadas geográficas do evento (formato JSON)
	Coordinates datatypes.JSON `json:"coordinates" swaggertype:"object"`
	// Informações do grupo associado (opcional, formato JSON)
	Group *datatypes.JSON `json:"group" swaggertype:"object"`
	// Papel do usuário no evento (owner, admin, member, subscriber)
	MemberRole string `json:"memberRole" example:"subscriber"`
}
