package input

type GetNearbyCitiesFilters struct {
	Latitude  float64 `validate:"required,gte=-90,lte=90"`
	Longitude float64 `validate:"required,gte=-180,lte=180"`
}
