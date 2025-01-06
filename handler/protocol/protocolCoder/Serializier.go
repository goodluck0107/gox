package protocolCoder

import (
	"gitee.com/andyxt/gox/handler/protocol"
)

type Serializier interface {
	Serialize(protocol.Protocol) []byte
	Deserialize(b []byte) (bool, protocol.Protocol)
}
