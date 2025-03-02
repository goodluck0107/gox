package code

import (
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/handler/message"
)

// message.IMessage ---> []byte
type MessageEncoder struct {
}

func NewMessageEncoder() (this *MessageEncoder) {
	this = new(MessageEncoder)
	return
}

func (encoder *MessageEncoder) ExceptionCaught(ctx channel.ChannelContext, err error) {
	//	logger.Debug("MessageEncoder ExceptionCaught")
}

func (encoder *MessageEncoder) Write(ctx channel.ChannelContext, e interface{}) interface{} {
	//	logger.Debug("MessageEncoder Write")
	msg := e.(message.IMessage)
	buf, err := msg.Encode()
	if err != nil {
		logger.Error("MessageEncoder Write err=", err)
		return nil
	}
	return buf
}

func (encoder *MessageEncoder) Close(ctx channel.ChannelContext) {
	//	logger.Debug("MessageEncoder Close")
}
