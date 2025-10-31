package input

type GetPendingInvitesFilters struct {
	Page  uint `validate:"omitempty,number,gt=0"`
	Limit uint `validate:"omitempty,number,gt=0,max=20"`
}
