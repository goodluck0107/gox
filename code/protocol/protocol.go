package protocol

type Protocol interface {
	Decode(b []byte) error
	Encode() ([]byte, error)
	GetSeqID() uint32
	GetMsgID() uint16
	GetMsgData() []byte
}

type ProtocolFactory interface {
	GetProtocol(buf []byte) (valid bool, ret Protocol)
}
