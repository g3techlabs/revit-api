package input

import (
	"time"

	"github.com/g3techlabs/revit-api/src/core/event/entities"
	"github.com/g3techlabs/revit-api/src/db/models"
	"gorm.io/gorm"
)

// CreateEventInput representa os dados para criação de um novo evento
// @Description Dados necessários para criar um novo evento no sistema
type CreateEventInput struct {
	// Nome do evento
	Name string `json:"name" validate:"required" example:"Pedal Noturno"`
	// Descrição do evento
	Description string `json:"description" validate:"required" example:"Passeio ciclístico noturno pela cidade"`
	// Data e hora do evento no formato RFC3339 (2006-01-02T15:04:05Z07:00)
	Date string `json:"date" validate:"required,datetime=2006-01-02T15:04:05Z07:00" example:"2024-12-25T20:00:00-03:00"`
	// Coordenadas geográficas da localização do evento
	Location entities.Coordinates `json:"location" validate:"required"`
	// Tipo de conteúdo da foto do evento (opcional: image/jpeg, image/png, image/webp)
	PhotoContentType *string `json:"photoContentType" validate:"omitempty,oneof=image/jpeg image/png image/webp" example:"image/jpeg"`
	// ID da cidade onde o evento ocorrerá
	CityID uint `json:"cityId" validate:"required,number,gt=0" example:"1"`
	// Visibilidade do evento (public ou private)
	Visibility string `json:"visibility" validate:"required,oneof=public private" example:"public"`
	// ID do grupo associado ao evento (opcional)
	GroupID *uint `json:"groupId" validate:"omitempty,number,gt=0" example:"5"`
}

func (i *CreateEventInput) ToEventModel() *models.Event {
	dateTime, err := time.Parse("2006-01-02T15:04:05Z07:00", i.Date)
	if err != nil {
		return nil
	}

	var newVisibility uint
	switch i.Visibility {
	case "public":
		newVisibility = 1
	case "private":
		newVisibility = 2
	}

	return &models.Event{
		Name:         i.Name,
		Description:  i.Description,
		Date:         dateTime,
		CityID:       i.CityID,
		VisibilityID: newVisibility,
		GroupID:      i.GroupID,
		Location:     gorm.Expr("ST_SetSRID(ST_MakePoint(?, ?), 4326)", i.Location.Longitude, i.Location.Latitude),
	}
}
