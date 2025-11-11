package repository

func (r *GeoLocationRepository) GetUsersInRoute(key string) (ids []uint, err error) {
	stringIds, err := r.getUsersWithinRoute(key)
	if err != nil {
		return nil, err
	}

	ids = r.convertNearbyUsersToIDs(stringIds, 0)
	return
}
