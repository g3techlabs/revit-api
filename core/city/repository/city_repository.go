package repository

import (
	"fmt"
	"strings"

	"github.com/g3techlabs/revit-api/core/city/response"
	"github.com/g3techlabs/revit-api/db"
	"github.com/g3techlabs/revit-api/db/models"
	"gorm.io/gorm"
)

type CityRepository interface {
	GetCities(cityName string, page, limit uint) (*[]response.GetCityReponse, error)
}

type cityRepository struct {
	db *gorm.DB
}

func NewCityRepository() CityRepository {
	return &cityRepository{db: db.Db}
}

func (cr *cityRepository) GetCities(cityName string, page, limit uint) (*[]response.GetCityReponse, error) {
	var cities []response.GetCityReponse

	pattern := fmt.Sprintf("%%%s%%", strings.ToLower(cityName))

	query := cr.db.Model(&models.City{}).
		Select("city.id as city_id", "city.name AS city_name", "state.name AS state_name", "country.name AS country_name").
		Joins("JOIN state ON city.state_id = state.id").
		Joins("JOIN country ON state.country_id = country.id").
		Where("city.name LIKE ?", pattern)

	if limit > 0 {
		offset := 0
		if page > 0 {
			offset = int((page - 1) * limit)
		}
		query = query.Limit(int(limit)).Offset(offset)
	}

	if err := query.Scan(&cities).Error; err != nil {
		return nil, err
	}

	if cities == nil {
		empty := make([]response.GetCityReponse, 0)
		return &empty, nil
	}

	return &cities, nil
}
