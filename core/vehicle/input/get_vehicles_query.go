package input

type GetVehiclesParams struct {
	Limit    uint   `validate:"omitempty,number,gt=0,max=20"`
	Page     uint   `validate:"omitempty,number,gt=0"`
	Nickname string `validate:"omitempty"`
}
