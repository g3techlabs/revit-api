package response

type GetCityReponse struct {
	CityId      uint   `json:"cityId"`
	CityName    string `json:"cityName"`
	StateName   string `json:"stateName"`
	CountryName string `json:"countryName"`
}
