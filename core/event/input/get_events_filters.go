package input

type GetEventsFilters struct {
	Name       string   `validate:"omitempty"`
	FromDate   string   `validate:"omitempty,datetime=2006-01-02"`
	ToDate     string   `validate:"omitempty,datetime=2006-01-02"`
	Latitude   *float64 `validate:"required_with=Longitude,gte=-90,lte=90"`
	Longitude  *float64 `validate:"required_with=Latitude,gte=-180,lte=180"`
	CityID     uint     `validate:"omitempty,number,gt=0"`
	MemberType *string  `validate:"omitempty,oneof=owner admin member"`
	Visibility string   `validate:"omitempty,oneof=public private"`
	Limit      uint     `validate:"omitempty,number,gt=0"`
	Page       uint     `validate:"omitempty,number,gt=0"`
}
