package input

type AnswerPendingInvite struct {
	Answer string `json:"answer" validate:"required,oneof=accept reject"`
}
