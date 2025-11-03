package websocket

type MulticastMessage struct {
	TargetUserIDs []uint
	Payload       []byte
}
