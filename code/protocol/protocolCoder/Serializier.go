package protocolCoder

import (
	"gitee.com/andyxt/gox/code/protocol"
)

type Serializier interface {
	Serialize(protocol.Protocol) []byte
	Deserialize(b []byte) (bool, protocol.Protocol)
}
