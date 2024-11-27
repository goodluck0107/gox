package channelCommands

import (
	"gitee.com/andyxt/gox/extends"
	"gitee.com/andyxt/gox/handler/schedule"
	"gitee.com/andyxt/gox/service"

	"gitee.com/andyxt/gona/logger"
	"gitee.com/andyxt/gox/executor"
)

type ClientCommandInActive struct {
	mEventMaker schedule.IRoutineInboundEventMaker
	ChannelCtx  service.IChannelContext
}

func NewClientCommandInActive(mEventMaker schedule.IRoutineInboundEventMaker, ChannelCtx service.IChannelContext) (this *ClientCommandInActive) {
	this = new(ClientCommandInActive)
	this.mEventMaker = mEventMaker
	this.ChannelCtx = ChannelCtx
	return
}

func (event *ClientCommandInActive) Exec() {
	logger.Debug("ClientCommandInActive Exec", extends.ChannelContextToString(event.ChannelCtx))
	if !extends.HasUserInfo(event.ChannelCtx) {
		logger.Debug("ClientCommandInActive Exec: ChannelCtx is not IsInPool", extends.ChannelContextToString(event.ChannelCtx))
		return
	}
	executor.FireEvent(event.mEventMaker.MakeInActiveEvent(extends.UID(event.ChannelCtx), event.ChannelCtx))
}
