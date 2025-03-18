package message

type MessageType uint32

const (
	MessageTypePB     MessageType = 0 // ProtoBuffer
	MessageTypeJson   MessageType = 1 // Json
	MessageTypeCustom MessageType = 2 // Custom
)
