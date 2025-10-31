package input

type FriendshipRequestAnswer struct {
	Answer string `json:"answer" validate:"required,oneof=accept reject"`
}
