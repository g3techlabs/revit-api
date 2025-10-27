package input

type GetCitiesQuery struct {
	Name  string `validate:"required"`
	Page  uint   `validate:"omitempty,number,gt=0"`
	Limit uint   `validate:"omitempty,number,gt=0"`
}
