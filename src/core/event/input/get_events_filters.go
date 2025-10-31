package input

type GetEventsFilters struct {
	Name       string   `validate:"omitempty"`
	FromDate   string   `validate:"omitempty,datetime=2006-01-02"`
	ToDate     string   `validate:"omitempty,datetime=2006-01-02"`
	Latitude   *float64 `validate:"required_with=Longitude,omitempty,latitude"`
	Longitude  *float64 `validate:"required_with=Latitude,omitempty,longitude"`
	CityID     uint     `validate:"omitempty,number,gt=0"`
	MemberType *string  `validate:"omitempty,oneof=owner admin member"`
	Visibility string   `validate:"omitempty,oneof=public private"`
	Limit      uint     `validate:"omitempty,number,gt=0"`
	Page       uint     `validate:"omitempty,number,gt=0"`
}
