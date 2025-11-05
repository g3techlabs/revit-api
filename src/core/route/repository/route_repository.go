package repository

import (
	geoinput "github.com/g3techlabs/revit-api/src/core/geolocation/geo_input"
	"github.com/g3techlabs/revit-api/src/db"
	"github.com/g3techlabs/revit-api/src/db/models"
	"gorm.io/gorm"
)

type RouteRepository interface {
	CreateRoute(userId uint, startLocation, destination geoinput.Coordinates) (uint, error)
}

type routeRepository struct {
	db *gorm.DB
}

func NewRouteRepository() RouteRepository {
	return &routeRepository{
		db: db.Db,
	}
}

func (r *routeRepository) CreateRoute(userId uint, startLocation, destination geoinput.Coordinates) (uint, error) {
	route := models.Route{
		Destination: gorm.Expr("ST_SetSRID(ST_MakePoint(?, ?), 4326)", destination.Long, destination.Lat),
	}

	r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&route).Error; err != nil {
			return err
		}

		routeOwner := models.RouteParticipant{
			UserID:        userId,
			RouteID:       route.ID,
			StartLocation: gorm.Expr("ST_SetSRID(ST_MakePoint(?, ?), 4326)", startLocation.Long, startLocation.Lat),
			IsOwner:       true,
		}

		if err := tx.Create(&routeOwner).Error; err != nil {
			return err
		}

		return nil
	})

	return route.ID, nil
}
