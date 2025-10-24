package input

type UpdateGroup struct {
	Name        *string `json:"name" validate:"omitempty"`
	Description *string `json:"description" validate:"omitempty"`
	CityID      *uint   `json:"cityId" validate:"omitempty,number,gt=0"`
	StateID     *uint   `json:"stateId" validate:"omitempty,number,gt=0"`
	Visibility  *string `json:"visibility" validate:"omitempty,oneof=public private"`
}
