package input

type GetMembersInput struct {
	Nickname *string `validate:"omitempty,min=3,max=32"`
	Limit    uint    `validate:"required,number,gt=0,max=50"`
	Page     uint    `validate:"required,number,gt=0"`
}
