package handler

import (
	"gitee.com/andyxt/gox/code/protocol"

	"gitee.com/andyxt/gona/boot/channel"
	"gitee.com/andyxt/gox/internal/logger"
)

// []byte ---> protocol.Protocol
type ProtocolDecoder struct {
	protocolFactory protocol.ProtocolFactory
}

func NewProtocolDecoder(protocolFactory protocol.ProtocolFactory) (this *ProtocolDecoder) {
	this = new(ProtocolDecoder)
	this.protocolFactory = protocolFactory
	return
}

func (decoder *ProtocolDecoder) MessageReceived(ctx channel.ChannelContext, e interface{}) (ret interface{}, goonNext bool) {
	//	logger.Debug("MessageDecoder MessageReceived")
	byteSlice := e.([]byte)
	valid, msg := decoder.protocolFactory.GetProtocol(byteSlice)
	if !valid {
		logger.Error("关闭连接：", " 关闭原因：协议解析失败:IP=", ctx.RemoteAddr()) //, " , 协议号：", VersionId, UserId, AppId, MessageId)
		ctx.Close()
		return
	}
	ret = msg
	goonNext = true
	return
}

func (decoder *ProtocolDecoder) ChannelActive(ctx channel.ChannelContext) (goonNext bool) {
	//	logger.Debug("MessageDecoder ChannelActive")
	return true
}

func (decoder *ProtocolDecoder) ChannelInactive(ctx channel.ChannelContext) (goonNext bool) {
	//	logger.Debug("MessageDecoder ChannelInactive")
	return true
}

func (decoder *ProtocolDecoder) ExceptionCaught(ctx channel.ChannelContext, err error) {
	//	logger.Debug("MessageDecoder ExceptionCaught")
}
