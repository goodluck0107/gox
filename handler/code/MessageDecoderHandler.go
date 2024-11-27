package code

import (
	"gitee.com/andyxt/gox/message"

	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gona/logger"
)

// *buffer.ProtocolBuffer ---> UpBase
type MessageDecoderHandleOnRoutineSync struct {
	messageFactory message.IMessageFactory
}

func NewMessageDecoderHandleOnRoutineSync(messageFactory message.IMessageFactory) (this *MessageDecoderHandleOnRoutineSync) {
	this = new(MessageDecoderHandleOnRoutineSync)
	this.messageFactory = messageFactory
	return
}

func (decoder *MessageDecoderHandleOnRoutineSync) MessageReceived(ctx channel.ChannelContext, e interface{}) (ret interface{}, goonNext bool) {
	//	logger.Debug("MessageDecoder MessageReceived")
	byteSlice := e.([]byte)
	valid, msg := decoder.messageFactory.GetMessage(byteSlice)
	if !valid {
		logger.Error("关闭连接：", " 关闭原因：协议解析失败:IP=", ctx.RemoteAddr()) //, " , 协议号：", VersionId, UserId, AppId, MessageId)
		ctx.Close()
		return
	}
	ret = msg
	goonNext = true
	return
}

func (decoder *MessageDecoderHandleOnRoutineSync) ChannelActive(ctx channel.ChannelContext) (goonNext bool) {
	//	logger.Debug("MessageDecoder ChannelActive")
	return true
}
func (decoder *MessageDecoderHandleOnRoutineSync) ChannelInactive(ctx channel.ChannelContext) (goonNext bool) {
	//	logger.Debug("MessageDecoder ChannelInactive")
	return true
}

func (decoder *MessageDecoderHandleOnRoutineSync) ExceptionCaught(ctx channel.ChannelContext, err error) {
	//	logger.Debug("MessageDecoder ExceptionCaught")
}
