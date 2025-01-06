package protocol

type Protocol interface {
	Decode(e interface{}) (valid bool)
	Encode() (ret interface{})
}

type ProtocolFactory interface {
	GetProtocol(buf []byte) (valid bool, ret Protocol)
}
