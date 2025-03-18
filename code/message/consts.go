package message

type MessageType uint32

const (
	ProtoTypePB   MessageType = 0 // ProtoBuffer
	ProtoTypeBN   MessageType = 1 // Binary
	ProtoTypeJson MessageType = 2 // Json
)
