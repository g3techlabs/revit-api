package input

type GetGroupsQuery struct {
	Name    string `validate:"omitempty"`
	CityId  uint   `validate:"omitempty,number,gt=0"`
	StateId uint   `validate:"omitempty,number,gt=0"`
	Member  bool   `validate:"omitempty"`
	Limit   uint   `validate:"omitempty,number,gt=0,max=20"`
	Page    uint   `validate:"omitempty,number,gt=0"`
}
