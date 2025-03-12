package channelCommands

import (
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/code/protocol"
	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/handler"
)

type ChannelInboundCmdMsgRecv struct {
	mEventMaker   handler.IRoutineInboundEventMaker
	ChannelCtx    service.IChannelContext
	e             interface{}
	mLoginMessage ILoginMessage
}

func NewChannelInboundCmdMsgRecv(
	mEventMaker handler.IRoutineInboundEventMaker, ChannelCtx service.IChannelContext, e interface{}, mLoginMessage ILoginMessage) (this *ChannelInboundCmdMsgRecv) {
	this = new(ChannelInboundCmdMsgRecv)
	this.mEventMaker = mEventMaker
	this.ChannelCtx = ChannelCtx
	this.e = e
	this.mLoginMessage = mLoginMessage
	return
}

func (event *ChannelInboundCmdMsgRecv) Exec() {
	logger.Debug("ChannelInboundCmdMsgRecv Exec !", extends.ChannelContextToString(event.ChannelCtx))
	buf, _ := event.e.(protocol.Protocol)
	if extends.HasUserInfo(event.ChannelCtx) { // 用户已经发送过登录协议
		logger.Debug("ChannelInboundCmdMsgRecv Exec executor.FireMessageReceivedEvent !", extends.ChannelContextToString(event.ChannelCtx))
		executor.FireEvent(event.mEventMaker.MakeMessageReceivedEvent(extends.UID(event.ChannelCtx), buf, event.ChannelCtx))
		return
	}
	logger.Debug("ChannelInboundCmdMsgRecv Exec ChannelCtx Not HasUserInfo !", extends.ChannelContextToString(event.ChannelCtx))
	if !event.mLoginMessage.IsLoginMessage(buf) {
		// 处理白名单消息
		if !event.mLoginMessage.IsWhiteMessage(buf) {
			logger.Debug("ChannelInboundCmdMsgRecv Exec First Message Is Not Login Message !", extends.ChannelContextToString(event.ChannelCtx), "关闭连接, 关闭原因：尚未通过验证的连接发送任何非登陆消息都认为是非法")
			extends.Close(event.ChannelCtx)
			return
		}
		logger.Debug("ChannelInboundCmdMsgRecv Exec First Message Is Not Login Message executor.FireMessageReceivedEvent !", extends.ChannelContextToString(event.ChannelCtx))
		executor.FireEvent(event.mEventMaker.MakeMessageReceivedEvent(0, buf, event.ChannelCtx))
		return
	}
	if !event.mLoginMessage.IsValid(buf) {
		logger.Debug("ChannelInboundCmdMsgRecv Exec Login Message Is Invalid !", extends.ChannelContextToString(event.ChannelCtx), "关闭连接, 关闭原因：发送非法登陆消息")
		extends.Close(event.ChannelCtx)
		return
	}
	uID := event.mLoginMessage.GetLoginUID(buf)
	lngType := event.mLoginMessage.GetLngType(buf)
	extends.PutInUserInfo(event.ChannelCtx, uID, lngType)
	logger.Debug("ChannelInboundCmdMsgRecv Exec Login Message executor.FireMessageReceivedEvent !", extends.ChannelContextToString(event.ChannelCtx))
	executor.FireEvent(event.mEventMaker.MakeMessageReceivedEvent(extends.UID(event.ChannelCtx), buf, event.ChannelCtx))
}
