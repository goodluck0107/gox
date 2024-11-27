package channelCommands

import (
	"gitee.com/andyxt/gox/message"
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/handler/protocol"
	"gitee.com/andyxt/gox/handler/schedule"
)

type ChannelInboundCmdMsgRecv struct {
	mEventMaker   schedule.IRoutineInboundEventMaker
	ChannelCtx    service.IChannelContext
	e             interface{}
	mLoginMessage ILoginMessage
}

func NewChannelInboundCmdMsgRecv(
	mEventMaker schedule.IRoutineInboundEventMaker, ChannelCtx service.IChannelContext, e interface{}, mLoginMessage ILoginMessage) (this *ChannelInboundCmdMsgRecv) {
	this = new(ChannelInboundCmdMsgRecv)
	this.mEventMaker = mEventMaker
	this.ChannelCtx = ChannelCtx
	this.e = e
	this.mLoginMessage = mLoginMessage
	return
}

func (event *ChannelInboundCmdMsgRecv) Exec() {
	logger.Debug("ChannelInboundCmdMsgRecv Exec !", extends.ChannelContextToString(event.ChannelCtx))
	buf, _ := event.e.(protocol.IProtocol)
	if extends.HasUserInfo(event.ChannelCtx) { // 用户已经发送过登录协议
		logger.Debug("ChannelInboundCmdMsgRecv Exec executor.FireMessageReceivedEvent !", extends.ChannelContextToString(event.ChannelCtx))
		executor.FireEvent(event.mEventMaker.MakeMessageReceivedEvent(extends.UID(event.ChannelCtx), buf, event.ChannelCtx))
		return
	}
	logger.Debug("ChannelInboundCmdMsgRecv Exec ChannelCtx Not HasUserInfo !", extends.ChannelContextToString(event.ChannelCtx))
	msgData, ok := buf.(*message.Message)
	if !ok {
		logger.Debug("ChannelInboundCmdMsgRecv Exec Message Is Invalid  !", extends.ChannelContextToString(event.ChannelCtx), "关闭连接, 关闭原因：无效消息")
		extends.Close(event.ChannelCtx)
		return
	}
	if !event.mLoginMessage.IsLoginMessage(msgData) {
		logger.Debug("ChannelInboundCmdMsgRecv Exec First Message Is Not Login Message !", extends.ChannelContextToString(event.ChannelCtx), "关闭连接, 关闭原因：尚未通过验证的连接发送任何非登陆消息都认为是非法")
		extends.Close(event.ChannelCtx)
		return
	}
	if !event.mLoginMessage.IsValid(msgData) {
		logger.Debug("ChannelInboundCmdMsgRecv Exec Login Message Is Invalid !", extends.ChannelContextToString(event.ChannelCtx), "关闭连接, 关闭原因：发送非法登陆消息,消息内容：", msgData.Data)
		extends.Close(event.ChannelCtx)
		return
	}
	uID := event.mLoginMessage.GetLoginUID(msgData)
	lngType := event.mLoginMessage.GetLngType(msgData)
	extends.PutInUserInfo(event.ChannelCtx, uID, lngType)
	logger.Debug("ChannelInboundCmdMsgRecv Exec Login Message executor.FireMessageReceivedEvent !", extends.ChannelContextToString(event.ChannelCtx))
	executor.FireEvent(event.mEventMaker.MakeMessageReceivedEvent(extends.UID(event.ChannelCtx), buf, event.ChannelCtx))
}
