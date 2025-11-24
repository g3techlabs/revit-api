package response

import "time"

// Vehicle representa um veículo retornado na listagem ou detalhamento
// @Description Informações completas de um veículo retornado na busca/listagem
type Vehicle struct {
	// ID do veículo
	ID uint `json:"id" example:"123"`
	// Apelido do veículo
	Nickname string `json:"nickname" example:"Ligeirinho"`
	// Marca do veículo
	Brand string `json:"brand" example:"Honda"`
	// Modelo do veículo
	Model string `json:"model" example:"Civic"`
	// Ano do veículo
	Year uint `json:"year" example:"1994"`
	// Versão do veículo (opcional)
	Version *string `json:"version" example:"Hatch EG8"`
	// URL da foto principal do veículo (opcional)
	MainPhotoUrl *string `json:"mainPhotoUrl" example:"https://example.com/vehicles/123/main.jpg"`
	// Data de criação do veículo
	CreatedAt time.Time `json:"createdAt" example:"2024-01-15T10:30:00Z"`
	// Lista de fotos adicionais do veículo
	Photos []Photo `json:"photos"`
}

// Photo representa uma foto adicional do veículo
// @Description Informações sobre uma foto adicional do veículo
type Photo struct {
	// ID da foto
	ID uint `json:"id" example:"456"`
	// URL da foto
	Url string `json:"reference" example:"https://example.com/vehicles/123/photos/456.jpg"`
}
