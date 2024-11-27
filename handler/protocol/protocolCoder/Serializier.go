package protocolCoder

import (
	"gitee.com/andyxt/gox/handler/protocol"
)

type Serializier interface {
	Serialize(protocol.IProtocol) []byte
	Deserialize(b []byte) (bool, protocol.IProtocol)
}
