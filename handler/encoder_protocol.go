package handler

import (
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/code/protocol"
)

// protocol.Protocol ---> []byte
type ProtocolEncoder struct {
}

func NewProtocolEncoder() (this *ProtocolEncoder) {
	this = new(ProtocolEncoder)
	return
}

func (encoder *ProtocolEncoder) ExceptionCaught(ctx channel.ChannelContext, err error) {
	//	logger.Debug("MessageEncoder ExceptionCaught")
}

func (encoder *ProtocolEncoder) Write(ctx channel.ChannelContext, e interface{}) interface{} {
	//	logger.Debug("MessageEncoder Write")
	proto := e.(protocol.Protocol)
	buf, err := proto.Encode()
	if err != nil {
		logger.Error("ProtocolEncoder Write err=", err)
		return nil
	}
	return buf
}

func (encoder *ProtocolEncoder) Close(ctx channel.ChannelContext) {
	//	logger.Debug("MessageEncoder Close")
}
