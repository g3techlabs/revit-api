package input

type GetGroupsQuery struct {
	Name       string `validate:"omitempty"`
	CityId     uint   `validate:"omitempty,number,gt=0"`
	StateId    uint   `validate:"omitempty,number,gt=0"`
	Visibility string `validate:"omitempty,oneof=public private"`
	Member     bool
	Limit      uint `validate:"omitempty,number,gt=0,max=20"`
	Page       uint `validate:"omitempty,number,gt=0"`
}
