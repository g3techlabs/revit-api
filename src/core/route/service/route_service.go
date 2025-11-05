package service

import (
	"github.com/g3techlabs/revit-api/src/core/geolocation/service"
	"github.com/g3techlabs/revit-api/src/core/route/input"
	"github.com/g3techlabs/revit-api/src/core/route/repository"
	"github.com/g3techlabs/revit-api/src/validation"
)

type IRouteService interface {
	CreateRoute(userId uint, data *input.CreateRouteInput) error
}

type RouteService struct {
	validator          validation.IValidator
	geoLocationService service.IGeoLocationService
	routeRepo          repository.RouteRepository
}

func NewRouteService(validator validation.IValidator, geoLocationService service.IGeoLocationService, routeRepo repository.RouteRepository) IRouteService {
	return &RouteService{validator: validator, routeRepo: routeRepo, geoLocationService: geoLocationService}
}
