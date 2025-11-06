package service

func (gls *GeoLocationService) CheckUsersAreOnline(userIds []uint) ([]bool, error) {
	return gls.geoLocationRepository.CheckUsersAreOnline(userIds)
}
