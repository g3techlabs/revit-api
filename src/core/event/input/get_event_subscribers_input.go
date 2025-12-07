package input

type GetEventSubscribersInput struct {
	Nickname *string `json:"nickname" validate:"omitempty,min=3,max=32"`
	Limit    uint    `json:"limit" validate:"required,number,gt=0,max=50"`
	Page     uint    `json:"page" validate:"required,number,gt=0"`
}
