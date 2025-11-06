package service

func (s *GeoLocationService) GetNearbyUsersToRouteInvite(userId uint, lat, long float64, page, pageSize int) ([]uint, error) {
	return s.geoLocationRepository.GetNearbyUsersToRouteInvite(userId, lat, long, page, pageSize)
}
