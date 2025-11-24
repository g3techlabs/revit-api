package input

type NicknameInput struct {
	Nickname string `json:"nickname" validate:"required,min=3,max=32,lowercase"`
}
