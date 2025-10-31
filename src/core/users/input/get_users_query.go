package input

type GetUsersQuery struct {
	Page     uint   `validate:"omitempty,number,gt=0"`
	Limit    uint   `validate:"omitempty,number,gt=0,max=20"`
	Nickname string `validate:"omitempty,min=3"`
}
