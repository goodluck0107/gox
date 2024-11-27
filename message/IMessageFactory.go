package message

import "gitee.com/andyxt/gox/handler/protocol"

type IMessageFactory interface {
	GetMessage(buf []byte) (valid bool, ret protocol.IProtocol)
}
