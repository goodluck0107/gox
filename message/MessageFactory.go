package message

import "gitee.com/andyxt/gox/handler/protocol"

type MessageFactory struct {
}

func NewMessageFactory() (this *MessageFactory) {
	this = new(MessageFactory)
	return
}

func (factory *MessageFactory) GetMessage(buf []byte) (valid bool, ret protocol.IProtocol) {
	msg := NewMessageWith(buf)
	if msg != nil {
		return true, msg
	}
	return false, nil
}
