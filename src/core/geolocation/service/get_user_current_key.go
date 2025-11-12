package service

func (s *GeoLocationService) GetUserCurrentKey(userId uint) (string, error) {
	return s.geoLocationRepository.GetUserStateGeoKey(userId)
}
