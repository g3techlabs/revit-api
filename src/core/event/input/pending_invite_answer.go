package input

type PendingInviteAnswer struct {
	Answer string `json:"answer" validate:"required,oneof=accept reject"`
}
