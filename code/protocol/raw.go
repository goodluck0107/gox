package protocol

import (
	"gitee.com/andyxt/gox/code/message"
)

func Raw(v message.IMessage) Protocol {
	return &rawMessage{v: v}
}

type rawMessage struct {
	v message.IMessage
}

func (bean *rawMessage) Decode(e []byte) error {
	return bean.v.Decode(e)
}

func (bean *rawMessage) Encode() ([]byte, error) {
	return bean.v.Encode()
}

func (bean *rawMessage) GetSeqID() uint32 {
	return 0
}

func (bean *rawMessage) GetMsgID() uint16 {
	return 0
}

func (bean *rawMessage) GetMsgData() []byte {
	return []byte{}
}
