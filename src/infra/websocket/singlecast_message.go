package websocket

type SingleCastMessage struct {
	TargetUserID uint
	Payload      []byte
}
