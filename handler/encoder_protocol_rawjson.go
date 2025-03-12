package handler

import (
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/code/message"
	"gitee.com/andyxt/gox/code/protocol"
)

func NewProtocolRawJsonEncoder() (this *ProtocolRawJsonEncoder) {
	this = new(ProtocolRawJsonEncoder)
	return
}

// json ---> []byte
type ProtocolRawJsonEncoder struct {
}

func (encoder *ProtocolRawJsonEncoder) ExceptionCaught(ctx channel.ChannelContext, err error) {
	//	logger.Debug("MessageEncoder ExceptionCaught")
}

func (encoder *ProtocolRawJsonEncoder) Write(ctx channel.ChannelContext, e interface{}) interface{} {
	//	logger.Debug("MessageEncoder Write")
	proto := protocol.Raw(message.Json(e))
	buf, err := proto.Encode()
	if err != nil {
		logger.Error("ProtocolRawJsonEncoder Write err=", err)
		return nil
	}
	return buf
}

func (encoder *ProtocolRawJsonEncoder) Close(ctx channel.ChannelContext) {
	//	logger.Debug("MessageEncoder Close")
}
