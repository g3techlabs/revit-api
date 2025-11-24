package response

// GetCityReponse representa uma cidade retornada na listagem
// @Description Informações de uma cidade retornada na busca/listagem
type GetCityReponse struct {
	// ID da cidade
	CityId uint `json:"cityId" example:"1"`
	// Nome da cidade
	CityName string `json:"cityName" example:"São Paulo"`
	// Nome do estado
	StateName string `json:"stateName" example:"SP"`
	// Nome do país
	CountryName string `json:"countryName" example:"Brasil"`
}
