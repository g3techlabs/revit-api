package input

type GetOnlineFriendsToInviteQuery struct {
	Limit uint `validate:"required,number,gt=0,lt=100"`
	Page  uint `validate:"required,number,gt=0"`
}
