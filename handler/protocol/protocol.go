package protocol

type Protocol interface {
	Decode(e interface{}) (valid bool)
	Encode() (ret interface{})
	GetSeqID() uint32
	GetMsgID() uint16
	GetMsgData() []byte
}

type ProtocolFactory interface {
	GetProtocol(buf []byte) (valid bool, ret Protocol)
}
