package input

type RequestFriendship struct {
	DestinataryId uint `json:"destinataryId" validate:"required,number,gt=0"`
}
