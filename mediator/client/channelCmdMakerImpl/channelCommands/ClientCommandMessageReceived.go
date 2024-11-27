package channelCommands

import (
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/handler/protocol"
	"gitee.com/andyxt/gox/handler/schedule"
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/executor"
)

type ClientCommandMessageReceived struct {
	mEventMaker schedule.IRoutineInboundEventMaker
	ChannelCtx  service.IChannelContext
	e           interface{}
}

func NewClientCommandMessageReceived(mEventMaker schedule.IRoutineInboundEventMaker, ChannelCtx service.IChannelContext, e interface{}) (this *ClientCommandMessageReceived) {
	this = new(ClientCommandMessageReceived)
	this.mEventMaker = mEventMaker
	this.ChannelCtx = ChannelCtx
	this.e = e
	return
}

func (event *ClientCommandMessageReceived) Exec() {
	logger.Debug("ClientCommandMessageReceived Exec", extends.ChannelContextToString(event.ChannelCtx))
	if event.ChannelCtx == nil || event.e == nil {
		return
	}
	if !extends.HasUserInfo(event.ChannelCtx) {
		logger.Debug("关闭连接：", " 关闭原因：MessageReceived but ChannelCtx is not InPool", extends.ChannelContextToString(event.ChannelCtx))
		event.ChannelCtx.Close()
		return
	}
	buf, _ := event.e.(protocol.IProtocol)
	executor.FireEvent(event.mEventMaker.MakeMessageReceivedEvent(extends.UID(event.ChannelCtx), buf, event.ChannelCtx))
}
