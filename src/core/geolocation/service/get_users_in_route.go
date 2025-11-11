package service

import "fmt"

func (s *GeoLocationService) GetUsersInRoute(routeId uint) ([]uint, error) {
	routeKey := fmt.Sprintf("route:%d", routeId)

	usersIds, err := s.geoLocationRepository.GetUsersInRoute(routeKey)
	if err != nil {
		s.logger.Errorf("Error retrieving Users in %s: %v", routeKey, err)
		return nil, err
	}

	return usersIds, nil
}
