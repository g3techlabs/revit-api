package repository

import (
	"fmt"
	"strings"

	"github.com/g3techlabs/revit-api/core/city/response"
	"github.com/g3techlabs/revit-api/db"
	"github.com/g3techlabs/revit-api/db/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CityRepository interface {
	GetCities(cityName string, page, limit uint) (*[]response.GetCityReponse, error)
	GetNearbyCities(latitude, longitude float64) (*[]response.GetCityReponse, error)
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
		Where("lower(city.name) LIKE ?", pattern).
		Order(clause.Expr{SQL: "similarity(city.name, ?) DESC", Vars: []interface{}{cityName}})

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

func (cr *cityRepository) GetNearbyCities(latitude, longitude float64) (*[]response.GetCityReponse, error) {
	var nearbyCities []response.GetCityReponse

	sql := `
		SELECT 
			city.id AS city_id,
			city.name AS city_name,
			state.name AS state_name,
			country.name AS country_name
		FROM city
		JOIN state ON city.state_id = state.id
		JOIN country ON state.country_id = country.id
		ORDER BY ST_Distance(location, ST_MakePoint(?, ?)::geography) ASC
		LIMIT 4
	`

	if err := cr.db.Raw(sql, latitude, longitude).Scan(&nearbyCities).Error; err != nil {
		return nil, err
	}

	if len(nearbyCities) == 0 {
		return &[]response.GetCityReponse{}, nil
	}

	return &nearbyCities, nil
}
