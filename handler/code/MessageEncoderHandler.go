package code

import (
	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gox/handler/protocol"
)

// DownBase ---> *buffer.ProtocolBuffer
type MessageEncoderHandleOnRoutineSync struct {
}

func NewMessageEncoderHandleOnRoutineSync() (this *MessageEncoderHandleOnRoutineSync) {
	this = new(MessageEncoderHandleOnRoutineSync)
	return
}

func (encoder *MessageEncoderHandleOnRoutineSync) ExceptionCaught(ctx channel.ChannelContext, err error) {
	//	logger.Debug("MessageEncoder ExceptionCaught")
}

func (encoder *MessageEncoderHandleOnRoutineSync) Write(ctx channel.ChannelContext, e interface{}) (ret interface{}) {
	//	logger.Debug("MessageEncoder Write")
	buf := e.(protocol.IProtocol)
	ret = buf.Encode()
	return
}

func (encoder *MessageEncoderHandleOnRoutineSync) Close(ctx channel.ChannelContext) {
	//	logger.Debug("MessageEncoder Close")
}
