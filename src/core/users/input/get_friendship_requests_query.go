package input

type GetFriendshipRequestsQuery struct {
	Page  uint `validate:"omitempty,number,gt=0"`
	Limit uint `validate:"omitempty,number,gt=0,max=20"`
}
