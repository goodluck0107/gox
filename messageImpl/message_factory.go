package messageImpl

import "gitee.com/andyxt/gox/code/protocol"

type MessageFactory struct {
}

func NewMessageFactory() (this *MessageFactory) {
	this = new(MessageFactory)
	return
}

func (factory *MessageFactory) GetProtocol(buf []byte) (valid bool, ret protocol.Protocol) {
	msg := NewMessageWith(buf)
	if msg != nil {
		return true, msg
	}
	return false, nil
}
