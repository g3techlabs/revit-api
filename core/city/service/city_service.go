package service

import (
	"github.com/g3techlabs/revit-api/core/city/input"
	"github.com/g3techlabs/revit-api/core/city/repository"
	"github.com/g3techlabs/revit-api/core/city/response"
	"github.com/g3techlabs/revit-api/validation"
)

type ICityService interface {
	GetCities(query *input.GetCitiesQuery) (*[]response.GetCityReponse, error)
}

type CityService struct {
	validator validation.IValidator
	cityRepo  repository.CityRepository
}

func NewCityService(validator validation.IValidator, cityRepo repository.CityRepository) ICityService {
	return &CityService{validator: validator, cityRepo: cityRepo}
}

func (cs *CityService) GetCities(query *input.GetCitiesQuery) (*[]response.GetCityReponse, error) {
	if err := cs.validator.Validate(query); err != nil {
		return nil, err
	}

	cities, err := cs.cityRepo.GetCities(query.Name, query.Page, query.Limit)
	if err != nil {
		return nil, err
	}

	return cities, nil
}
