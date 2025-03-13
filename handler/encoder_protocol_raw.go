package handler

import (
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gox/code/message"
	"gitee.com/andyxt/gox/code/protocol"
	"gitee.com/andyxt/gox/internal/logger"
)

func NewProtocolRawEncoder() (this *ProtocolRawEncoder) {
	this = new(ProtocolRawEncoder)
	return
}

// json ---> []byte
type ProtocolRawEncoder struct {
}

func (encoder *ProtocolRawEncoder) ExceptionCaught(ctx channel.ChannelContext, err error) {
	//	logger.Debug("MessageEncoder ExceptionCaught")
}

func (encoder *ProtocolRawEncoder) Write(ctx channel.ChannelContext, e interface{}) interface{} {
	//	logger.Debug("MessageEncoder Write")
	msg := e.(message.IMessage)
	proto := protocol.Raw(msg)
	buf, err := proto.Encode()
	if err != nil {
		logger.Error("ProtocolRawEncoder Write err=", err)
		return nil
	}
	return buf
}

func (encoder *ProtocolRawEncoder) Close(ctx channel.ChannelContext) {
	//	logger.Debug("MessageEncoder Close")
}
