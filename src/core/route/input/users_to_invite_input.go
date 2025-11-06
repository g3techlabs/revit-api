package input

type UsersToInviteInput struct {
	IdsToInvite []uint `json:"idsToInvite" validate:"required,min=1,dive,number,gt=0"`
}
